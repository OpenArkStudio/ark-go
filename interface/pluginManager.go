package ark

import (
	"encoding/xml"
	"errors"
	"path/filepath"
	"sync"

	"ark-go/base"
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

// ------------------- AFPluginManager -------------------
type AFPluginManager struct {
	busId          int    // bus id
	timestamp      int64  // loop timestamp
	pluginPath     string // the xxxPlugin.so filepath
	resPath        string // the resource filepath
	pluginConfPath string // plugin configuration filepath
	appName        string // app name
	logPath        string // log output path

	pluginNameList            map[string]bool           // plugin_name -> bool
	orderedPluginNameList     []string                  // ordered plugin names
	pluginLibList             map[string]*base.AFDynLib // dynamic libraries
	pluginInstanceList        map[string]AFIPlugin      // plugin instances
	moduleInstanceList        map[string]AFIModule      // module instances
	orderedModuleInstanceList []AFIModule               // ordered module instances

	moduleWithUpdateFuncList map[string]AFIModule // the list of modules who have the `update` function
}

func GetAFPluginManagerInstance() *AFPluginManager {
	once.Do(func() {
		afPluginManager = &AFPluginManager{
			timestamp:                 util.GetNowTime(),
			pluginNameList:            make(map[string]bool),
			orderedPluginNameList:     make([]string, 0),
			pluginLibList:             make(map[string]*base.AFDynLib),
			pluginInstanceList:        make(map[string]AFIPlugin),
			moduleInstanceList:        make(map[string]AFIModule),
			orderedModuleInstanceList: make([]AFIModule, 0),
			moduleWithUpdateFuncList:  make(map[string]AFIModule),
		}
	})

	return afPluginManager
}

// ------------------- public func -------------------

func (a *AFPluginManager) Start() error {
	funcMap := []func() error{
		a.init,
		a.postInit,
		a.checkConfig,
		a.preUpdate,
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
		a.preShut,
		a.shut,
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

func (a *AFPluginManager) FindModule(name string) AFIModule {
	return a.moduleInstanceList[name]
}

func (a *AFPluginManager) Register(plugin AFIPlugin) {
	a.register(plugin)
}

func (a *AFPluginManager) Deregister(name string) {
	a.deregister(a.findPlugin(name))
}

func (a *AFPluginManager) AddModule(moduleName string, modulePtr AFIModule) {
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

func (a *AFPluginManager) AddUpdateModule(module AFIModule) error {
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

	//if !strings.Contains(path, ".plugin") {
	//	fmt.Println("failed to SetPluginConf  :", path)
	//	return
	//}

	a.pluginConfPath = path
}

func (a *AFPluginManager) GetLogPath() string {
	return a.logPath
}

func (a *AFPluginManager) SetLogPath(path string) {
	a.logPath = path
}

// ------------------- private func -------------------

func (a *AFPluginManager) register(plugin AFIPlugin) {
	pluginName := plugin.GetPluginName()
	if a.findPlugin(pluginName) != nil {
		return
	}

	plugin.SetPluginManager(a)
	a.pluginInstanceList[pluginName] = plugin
	plugin.Install()
}

func (a *AFPluginManager) deregister(plugin AFIPlugin) {
	if plugin == nil {
		return
	}

	plugin.Uninstall()
	delete(a.pluginInstanceList, plugin.GetPluginName())
}

func (a *AFPluginManager) findPlugin(pluginName string) AFIPlugin {
	return a.pluginInstanceList[pluginName]
}

func (a *AFPluginManager) init() error {
	// load plugin configuration
	if err := a.loadPluginConf(); err != nil {
		return err
	}

	// load plugin dynamic libraries
	for _, pluginName := range a.orderedPluginNameList {
		if err := a.loadPluginLibrary(pluginName); err != nil {
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

func (a *AFPluginManager) checkConfig() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.CheckConfig()
	}

	return nil
}

func (a *AFPluginManager) preUpdate() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PreUpdate()
	}

	return nil
}

func (a *AFPluginManager) preShut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PreShut()
	}

	return nil
}

func (a *AFPluginManager) shut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.Shut()
	}

	for pluginName, _ := range a.pluginNameList {
		a.unLoadPluginLibrary(pluginName)
	}

	for _, lib := range a.pluginLibList {
		lib.Unload()
	}

	return nil
}

func (a *AFPluginManager) loadPluginConf() error {
	cfg := &pluginConf{}
	data, err := util.GetBytes(a.pluginConfPath)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(data, cfg); err != nil {
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

func (a *AFPluginManager) loadPluginLibrary(pluginName string) error {
	if _, isExist := a.pluginLibList[pluginName]; isExist {
		// the plugin is already exist
		return nil
	}

	pDynLib := base.NewAFDynLib(pluginName)
	if err := pDynLib.Load(filepath.Join(a.pluginPath, pluginName+".so")); err != nil {
		return err
	}

	entryFunc, isOK := pDynLib.GetSymbol(entryPoint).(func(*AFPluginManager))
	if !isOK {
		// TODO： add log
		return errors.New("failed to get entryPoint")
	}

	entryFunc(a)

	return nil
}

func (a *AFPluginManager) unLoadPluginLibrary(pluginName string) error {
	pDynLib, isExist := a.pluginLibList[pluginName]
	if !isExist {
		return errors.New("plugin ` " + pluginName + " ` is absent")
	}

	exitFunc, isOK := pDynLib.GetSymbol(exitPoint).(func(*AFPluginManager))
	if !isOK {
		// TODO： add log
		return errors.New("")
	}

	exitFunc(a)

	return nil
}
