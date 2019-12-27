package model

type AFPlugin interface {
	GetPluginVersion() int
	GetPluginName() string
	Install()
	Uninstall()
	GetPluginManager() *AFPluginManager
	SetPluginManager(manager *AFPluginManager)
}

//type AFConsulPlugin struct {
//	modulesInPlugin map[string]*AFModule
//	manager         *AFPluginManager
//}
//
//type AFMysqlPlugin struct {
//	AFPlugin
//	modulesInPlugin map[string]*AFModule
//	manager         *AFPluginManager
//}
//
//func (plugin *AFConsulPlugin) GetPluginVersion() int {
//	return 1
//}
//
//func (plugin *AFConsulPlugin) GetPluginName() string {
//	return "consulPlugin"
//}
//
//func (plugin *AFConsulPlugin) Install() {
//	// TODO
//}
//
//func (plugin *AFConsulPlugin) Uninstall() {
//	// TODO
//}
//
//func (plugin *AFConsulPlugin) GetPluginManager() *AFPluginManager {
//	return plugin.manager
//}
//
//func (plugin *AFConsulPlugin) SetPluginManager(manager *AFPluginManager) {
//	plugin.manager = manager
//}
