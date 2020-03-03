package consulSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFConsulPlugin)(nil))

type AFConsulPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFConsulPlugin {
	return &AFConsulPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (consulPlugin *AFConsulPlugin) Install() {
	ark.RegisterModule(consulModuleType, &consulPlugin.AFCPlugin, consulModuleUpdate)
}

func (consulPlugin *AFConsulPlugin) Uninstall() {

	ark.DeregisterModule(consulModuleName, &consulPlugin.AFCPlugin)
}

func (consulPlugin *AFConsulPlugin) GetPluginName() string {
	return PluginName
}