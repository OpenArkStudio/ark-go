package consulPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule"

	_ "github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule/src"
)

var PluginName = ark.GetName((*AFConsulPlugin)(nil))

type AFConsulPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFConsulPlugin {
	return &AFConsulPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (consulPlugin *AFConsulPlugin) Install() {
	consulPlugin.AFCPlugin.RegisterModule(consulModule.ModuleType, consulModule.ModuleUpdate)
}

func (consulPlugin *AFConsulPlugin) Uninstall() {
	consulPlugin.AFCPlugin.DeregisterModule(consulModule.ModuleName)
}

func (consulPlugin *AFConsulPlugin) GetPluginName() string {
	return PluginName
}
