package logPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"

	_ "github.com/ArkNX/ark-go/plugin/logPlugin/logModule/src"
)

var PluginName = ark.GetName((*AFLogPlugin)(nil))

type AFLogPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFLogPlugin {
	return &AFLogPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (logPlugin *AFLogPlugin) Install() {
	logPlugin.AFCPlugin.RegisterModule(logModule.ModuleType, logModule.ModuleUpdate)
}

func (logPlugin *AFLogPlugin) Uninstall() {
	logPlugin.AFCPlugin.DeregisterModule(logModule.ModuleName)
}

func (logPlugin *AFLogPlugin) GetPluginName() string {
	return PluginName
}
