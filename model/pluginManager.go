package model

import (
	"encoding/xml"
	"errors"
	"reflect"
	"strings"
	"sync"

	. "ark-go/common"
	"ark-go/util"
)

var (
	once            sync.Once
	afPluginManager *AFPluginManager
)

const (
	entryPoint = "DllEntryPlugin"
	exitPoint  = "DllExitPlugin"
)

type DynEntryPluginFunc func(*AFPluginManager)
type DynExitPluginFunc func(*AFPluginManager)

type pluginConf struct {
	XMLName xml.Name `xml:"xml"`
	Plugins *plugins `xml:"plugins"`
	Res     *res     `xml:"res"`
}

type plugins struct {
	XMLName xml.Name  `xml:"plugins"`
	Path    string    `xml:"path,attr"`
	Plugin  []*plugin `xml:"plugin"`
}

type plugin struct {
	XMLName xml.Name `xml:"plugin"`
	Name    string   `xml:"name,attr"`
}

type res struct {
	XMLName xml.Name `xml:"res"`
	Path    string   `xml:"path,attr"`
}

type AFPluginManager struct {
	busId          int    // bus id
	timestamp      int64  // loop timestamp
	pluginPath     string // the plugin.so filepath
	resPath        string // the resource filepath
	pluginConfPath string // plugin configuration filepath
	appName        string // app name
	logPath        string // log output path

	pluginNameList            map[string]bool      // plugin_name -> bool
	orderedPluginNameList     []string             // ordered plugin names
	pluginLibList             map[string]*AFDynLib // dynamic libraries
	pluginInstanceList        map[string]AFPlugin  // plugin instances
	moduleInstanceList        map[string]*AFModule // module instances
	orderedModuleInstanceList []*AFModule          // ordered module instances

	moduleWithUpdateFuncList map[string]*AFModule // the list of modules who have the `update` function
}

func GetAFPluginManagerInstance() *AFPluginManager {
	if afPluginManager == nil {
		once.Do(func() {
			afPluginManager = &AFPluginManager{
				timestamp:                 util.GetNowTime(),
				pluginNameList:            make(map[string]bool),
				orderedPluginNameList:     make([]string, 0),
				pluginLibList:             make(map[string]*AFDynLib),
				pluginInstanceList:        make(map[string]AFPlugin),
				moduleInstanceList:        make(map[string]*AFModule),
				orderedModuleInstanceList: make([]*AFModule, 0),
				moduleWithUpdateFuncList:  make(map[string]*AFModule),
			}
		})
	}
	return afPluginManager
}

// ------------------- public func -------------------

func (a *AFPluginManager) Start() error {
	funcMap := []func() error{
		a.init,
		a.postInit,
		a.CheckConfig,
		a.PreUpdate,
	}

	for _, function := range funcMap {
		if err := function(); err != nil {
			return err
		}
	}

	return nil
}

func (a *AFPluginManager) Stop() error {
	funcMap := []func() error{
		a.PreShut,
		a.Shut,
	}

	for _, function := range funcMap {
		function()
	}

	return nil
}

func (a *AFPluginManager) Update() error {
	a.timestamp = util.GetNowTime()

	for _, moduleWithUpdateFunc := range a.moduleWithUpdateFuncList {
		if moduleWithUpdateFunc == nil {
			continue
		}

		moduleWithUpdateFunc.Update()
	}
	return nil
}

func (a *AFPluginManager) FindModule(t reflect.Type) *AFModule {
	return a.moduleInstanceList[t.String()]
}

func (a *AFPluginManager) Register(plugin AFPlugin) {
	a.register(plugin)
}

func (a *AFPluginManager) Deregister(t reflect.Type) {
	a.deregister(a.findPlugin(t.String()))
}

func (a *AFPluginManager) AddModule(moduleName string, modulePtr *AFModule) {
	if modulePtr == nil {
		return
	}

	a.moduleInstanceList[moduleName] = modulePtr
	a.orderedModuleInstanceList = append(a.orderedModuleInstanceList, modulePtr)
}

func (a *AFPluginManager) RemoveModule(moduleName string) {
	module, isExist := a.moduleInstanceList[moduleName]
	if !isExist {
		return
	}

	delete(a.moduleInstanceList, module.GetName())

	index := -1
	for tempIndex, tempModule := range a.orderedModuleInstanceList {
		if module == tempModule {
			index = tempIndex
			break
		}
	}

	length := len(a.orderedModuleInstanceList)
	if index != -1 {
		switch index {
		case 0:
			a.orderedModuleInstanceList = a.orderedModuleInstanceList[1:]
		case length:
			a.orderedModuleInstanceList = a.orderedModuleInstanceList[:length-1]
		default:
			a.orderedModuleInstanceList = append(a.orderedModuleInstanceList[:index], a.orderedModuleInstanceList[index+1:]...)
		}
	}
}

func (a *AFPluginManager) AddUpdateModule(module *AFModule) error {
	if module == nil {
		return errors.New("update module to add is nil")
	}

	a.moduleWithUpdateFuncList[module.GetName()] = module
	return nil
}

func (a *AFPluginManager) RemoveUpdateModule(moduleName string) {
	delete(a.moduleWithUpdateFuncList, moduleName)
}

func (a *AFPluginManager) GetNowTime() int64 {
	return a.timestamp
}

func (a *AFPluginManager) GetBusID() int {
	return a.busId
}

func (a *AFPluginManager) SetBusID(id int) {
	a.busId = id
}

func (a *AFPluginManager) GetAppName() string {
	return a.appName
}

func (a *AFPluginManager) SetAppName(appName string) {
	a.appName = appName
}

func (a *AFPluginManager) GetResPath() string {
	return a.resPath
}

func (a *AFPluginManager) SetPluginConf(path string) {
	if path == "" {
		return
	}

	if !strings.Contains(path, ".plugin") {
		return
	}

	a.pluginConfPath = path
}

func (a *AFPluginManager) GetLogPath() string {
	return a.logPath
}

func (a *AFPluginManager) SetLogPath(path string) {
	a.logPath = path
}

// ------------------- private func -------------------

func (a *AFPluginManager) register(plugin AFPlugin) {
	pluginName := plugin.GetPluginName()
	if a.findPlugin(pluginName) != nil {
		return
	}

	plugin.SetPluginManager(a)
	a.pluginInstanceList[pluginName] = plugin
	plugin.Install()
}

func (a *AFPluginManager) deregister(plugin AFPlugin) {
	if plugin == nil {
		return
	}

	plugin.Uninstall()

	delete(a.pluginInstanceList, plugin.GetPluginName())
}

func (a *AFPluginManager) findPlugin(pluginName string) AFPlugin {
	return a.pluginInstanceList[pluginName]
}

func (a *AFPluginManager) init() error {
	// load plugin configuration
	if err := a.LoadPluginConf(); err != nil {
		return err
	}

	// load plugin dynamic libraries
	for _, pluginName := range a.orderedPluginNameList {
		if err := a.LoadPluginLibrary(pluginName); err != nil {
			return err
		}
	}

	// initialize all modules
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.Init()
	}

	return nil
}

func (a *AFPluginManager) postInit() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PostInit()
	}

	return nil
}

func (a *AFPluginManager) CheckConfig() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.CheckConfig()
	}

	return nil
}

func (a *AFPluginManager) PreUpdate() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PreUpdate()
	}

	return nil
}

func (a *AFPluginManager) PreShut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PreShut()
	}

	return nil
}

func (a *AFPluginManager) Shut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.Shut()
	}

	for pluginName, _ := range a.pluginNameList {
		a.UnLoadPluginLibrary(pluginName)
	}

	for _, lib := range a.pluginLibList {
		lib.Unload()
	}

	return nil
}

func (a *AFPluginManager) LoadPluginConf() error {
	cfg := &pluginConf{}
	if err := xml.Unmarshal(nil, cfg); err != nil {
		return err
	}

	// pre check
	if cfg.Plugins.Path == "" {
		return errors.New("plugins path is empty")
	}

	if cfg.Res.Path == "" {
		return errors.New("res path is empty")
	}

	a.pluginPath = cfg.Plugins.Path
	a.resPath = cfg.Res.Path
	for _, plugin := range cfg.Plugins.Plugin {
		a.pluginNameList[plugin.Name] = true
		a.orderedPluginNameList = append(a.orderedPluginNameList, plugin.Name)
	}

	return nil
}

func (a *AFPluginManager) LoadPluginLibrary(pluginName string) error {
	pDynLib, isExist := a.pluginLibList[pluginName]
	if !isExist {
		return errors.New("plugin ` " + pluginName + " ` is absent")
	}

	entryFunc, isOK := pDynLib.GetSymbol(entryPoint).(DynEntryPluginFunc)
	if !isOK {
		// TODO： add log
		return errors.New("")
	}

	entryFunc(a)

	return nil
}

func (a *AFPluginManager) UnLoadPluginLibrary(pluginName string) error {
	pDynLib, isExist := a.pluginLibList[pluginName]
	if !isExist {
		return errors.New("plugin ` " + pluginName + " ` is absent")
	}

	exitFunc, isOK := pDynLib.GetSymbol(exitPoint).(DynExitPluginFunc)
	if !isOK {
		// TODO： add log
		return errors.New("")
	}

	exitFunc(a)

	return nil
}
