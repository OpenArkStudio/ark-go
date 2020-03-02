package ark

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
