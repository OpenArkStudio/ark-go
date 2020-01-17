package model

type AFModule struct{}

func (module *AFModule) Init() error                        { return nil }
func (module *AFModule) PostInit() error                    { return nil }
func (module *AFModule) CheckConfig() error                 { return nil }
func (module *AFModule) PreUpdate() error                   { return nil }
func (module *AFModule) Update() error                      { return nil }
func (module *AFModule) PreShut() error                     { return nil }
func (module *AFModule) Shut() error                        { return nil }
func (module *AFModule) GetPluginManager() *AFPluginManager { return nil }

// Do nothing in the module interface
func (module *AFModule) SetPluginManager(manager *AFPluginManager) {}

// Do nothing in the module interface
func (module *AFModule) GetName() string     { return "" }
func (module *AFModule) SetName(name string) {}

//type AFTestModule struct {
//	AFModule
//	// other data
//}

//func test() {
//	module := &AFTestModule{}
//	module.Init()
//}
