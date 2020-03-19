package httpPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule"

	_ "github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule/src"
)

var PluginName = ark.GetName((*AFHttpPlugin)(nil))

type AFHttpPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFHttpPlugin {
	return &AFHttpPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (httpPlugin *AFHttpPlugin) Install() {
	httpPlugin.AFCPlugin.RegisterModule(httpServerModule.ModuleType, httpServerModule.ModuleUpdate)
}

func (httpPlugin *AFHttpPlugin) Uninstall() {
	httpPlugin.AFCPlugin.DeregisterModule(httpServerModule.ModuleName)
}

func (httpPlugin *AFHttpPlugin) GetPluginName() string {
	return PluginName
}
