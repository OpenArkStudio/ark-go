package httpSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFHttpPlugin)(nil))

type AFHttpPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFHttpPlugin {
	return &AFHttpPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (httpPlugin *AFHttpPlugin) Install() {
	ark.RegisterModule(httpServerModuleType, &httpPlugin.AFCPlugin, httpServerModuleUpdate)
}

func (httpPlugin *AFHttpPlugin) Uninstall() {
	ark.DeregisterModule(httpServerModuleName, &httpPlugin.AFCPlugin)
}

func (httpPlugin *AFHttpPlugin) GetPluginName() string {
	return PluginName
}
