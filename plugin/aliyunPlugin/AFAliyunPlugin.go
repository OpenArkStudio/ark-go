package aliyunPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/aliyunPlugin/ossModule"
	_ "github.com/ArkNX/ark-go/plugin/aliyunPlugin/ossModule/src"
)

var PluginName = ark.GetName((*AFAliyunPlugin)(nil))

type AFAliyunPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFAliyunPlugin {
	return &AFAliyunPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (aliyunPlugin *AFAliyunPlugin) Install() {
	aliyunPlugin.AFCPlugin.RegisterModule(ossModule.ModuleType, ossModule.ModuleUpdate)
}

func (aliyunPlugin *AFAliyunPlugin) Uninstall() {
	aliyunPlugin.AFCPlugin.DeregisterModule(ossModule.ModuleName)
}

func (aliyunPlugin *AFAliyunPlugin) GetPluginName() string {
	return PluginName
}
