package ark

type AFIModule interface {
	Init() error
	PostInit() error
	CheckConfig() error
	PreUpdate() error
	Update() error
	PreShut() error
	Shut() error
	GetPluginManager() *AFPluginManager
	SetPluginManager(manager *AFPluginManager)
	GetName() string
	SetName(name string)
}

// ------------------- AFIModule implement -------------------
// ------------------- Eclectic solution for c++ macro -------------------
type AFCModule struct {
	pluginManager *AFPluginManager
	name          string
}

func (module *AFCModule) Init() error        { return nil }
func (module *AFCModule) PostInit() error    { return nil }
func (module *AFCModule) CheckConfig() error { return nil }
func (module *AFCModule) PreUpdate() error   { return nil }
func (module *AFCModule) Update() error      { return nil }
func (module *AFCModule) PreShut() error     { return nil }
func (module *AFCModule) Shut() error        { return nil }

// Do nothing in the module interface
func (module *AFCModule) GetName() string {
	return module.name
}

func (module *AFCModule) SetName(name string) {
	module.name = name
}

func (module *AFCModule) GetPluginManager() *AFPluginManager {
	return module.pluginManager
}

// Do nothing in the module interface
func (module *AFCModule) SetPluginManager(manager *AFPluginManager) {
	if manager != nil {
		module.pluginManager = manager
	}
}
