package aliyunSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFAliyunPlugin)(nil))

type AFAliyunPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFAliyunPlugin {
	return &AFAliyunPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (aliyunPlugin *AFAliyunPlugin) Install() {
	ark.RegisterModule(ossModuleType, &aliyunPlugin.AFCPlugin, ossModuleUpdate)
}

func (aliyunPlugin *AFAliyunPlugin) Uninstall() {

	ark.DeregisterModule(ossModuleName, &aliyunPlugin.AFCPlugin)
}

func (aliyunPlugin *AFAliyunPlugin) GetPluginName() string {
	return PluginName
}
