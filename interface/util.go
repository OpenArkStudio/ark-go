package ark

import (
	"log"
	"path/filepath"
	"reflect"
)

func RegisterModule(t reflect.Type, p *AFCPlugin, update bool) {
	pRegModule, ok := reflect.New(t).Interface().(AFIModule)
	if !ok {
		log.Fatalf("type %v should be a AFIModule\n", t)
	}
	pRegModuleName := filepath.Join(t.PkgPath(), t.Name())

	pluginManager := GetAFPluginManagerInstance()
	pRegModule.SetPluginManager(pluginManager)
	pRegModule.SetName(pRegModuleName)
	pluginManager.AddModule(pRegModuleName, pRegModule)
	p.Modules[pRegModuleName] = pRegModule

	if update {
		pluginManager.AddUpdateModule(pRegModule)
	}
}

func DeregisterModule(name string, p *AFCPlugin) {
	pluginManager := GetAFPluginManagerInstance()
	if pluginManager.FindModule(name) == nil {
		return
	}
	pluginManager.RemoveModule(name)
	pluginManager.RemoveUpdateModule(name)
	delete(p.Modules, name)
}
