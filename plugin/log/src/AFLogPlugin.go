package logSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFLogPlugin)(nil))

type AFLogPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFLogPlugin {
	return &AFLogPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (logPlugin *AFLogPlugin) Install() {
	ark.RegisterModule(logModuleType, &logPlugin.AFCPlugin, logModuleUpdate)
}

func (logPlugin *AFLogPlugin) Uninstall() {
	ark.DeregisterModule(logModuleName, &logPlugin.AFCPlugin)
}

func (logPlugin *AFLogPlugin) GetPluginName() string {
	return PluginName
}
