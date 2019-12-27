package model

type AFModule struct {
	manager *AFPluginManager
	name    string
}

func (module *AFModule) Init() error                               { return nil }
func (module *AFModule) PostInit() error                           { return nil }
func (module *AFModule) CheckConfig() error                        { return nil }
func (module *AFModule) PreUpdate() error                          { return nil }
func (module *AFModule) Update() error                             { return nil }
func (module *AFModule) PreShut() error                            { return nil }
func (module *AFModule) Shut() error                               { return nil }
func (module *AFModule) GetPluginManager() *AFPluginManager        { return module.manager }
func (module *AFModule) SetPluginManager(manager *AFPluginManager) { module.manager = manager }
func (module *AFModule) GetName() string                           { return module.name }
func (module *AFModule) SetName(name string)                       { module.name = name }

//type AFTestModule struct {
//	AFModule
//	// other data
//}

//func test() {
//	module := &AFTestModule{}
//	module.Init()
//}
