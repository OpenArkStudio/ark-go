package ark

import (
	"log"
	"path/filepath"
	"reflect"
)

type AFIPlugin interface {
	GetPluginVersion() int
	GetPluginName() string
	Install()
	Uninstall()
	GetPluginManager() *AFPluginManager
	SetPluginManager(manager *AFPluginManager)
}

// ------------------- AFIPlugin implement -------------------
type AFCPlugin struct {
	Modules       map[string]AFIModule
	pluginManager *AFPluginManager
}

func NewAFCPlugin() AFCPlugin {
	return AFCPlugin{
		Modules:       make(map[string]AFIModule),
		pluginManager: nil,
	}
}

func (plugin *AFCPlugin) GetPluginVersion() int { return 0 }
func (plugin *AFCPlugin) GetPluginName() string { return "" }
func (plugin *AFCPlugin) Install()              {}
func (plugin *AFCPlugin) Uninstall()            {}
func (plugin *AFCPlugin) GetPluginManager() *AFPluginManager {
	return plugin.pluginManager
}
func (plugin *AFCPlugin) SetPluginManager(p *AFPluginManager) {
	plugin.pluginManager = p
}

func (plugin *AFCPlugin) RegisterModule(t reflect.Type, update string) {
	pRegModule, ok := reflect.New(t).Interface().(AFIModule)
	if !ok {
		log.Fatalf("type %v should be a AFIModule\n", t)
	}
	pRegModuleName := filepath.Join(t.PkgPath(), t.Name())

	pluginManager := GetAFPluginManagerInstance()
	pRegModule.SetPluginManager(pluginManager)
	pRegModule.SetName(pRegModuleName)
	pluginManager.AddModule(pRegModuleName, pRegModule)
	plugin.Modules[pRegModuleName] = pRegModule

	if update != afcModuleUpdate {
		pluginManager.AddUpdateModule(pRegModule)
	}
}

func (plugin *AFCPlugin) DeregisterModule(name string) {
	pluginManager := GetAFPluginManagerInstance()
	if pluginManager.FindModule(name) == nil {
		return
	}
	pluginManager.RemoveModule(name)
	pluginManager.RemoveUpdateModule(name)
	delete(plugin.Modules, name)
}
