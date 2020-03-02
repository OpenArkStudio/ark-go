package logSrc

import (
	"ark-go/interface"
	"ark-go/util"
)

var LogPluginName = util.GetName((*AFLogPlugin)(nil))

type AFLogPlugin struct {
	ark.AFCPlugin
}

func NewAFLogPlugin() *AFLogPlugin {
	return &AFLogPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (logPlugin *AFLogPlugin) Install() {
	ark.RegisterModule(logModuleType, &logPlugin.AFCPlugin, logModuleUpdate)
}

func (logPlugin *AFLogPlugin) Uninstall() {
	ark.DeregisterModule(logModuleName, &logPlugin.AFCPlugin)
}

func (logPlugin *AFLogPlugin) GetPluginName() string {
	return LogPluginName
}
