package kernelSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFKernelPlugin)(nil))

type AFKernelPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFKernelPlugin {
	return &AFKernelPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (kernelPlugin *AFKernelPlugin) Install() {
	kernelPlugin.AFCPlugin.RegisterModule(MetaClassModuleType, MetaClassModuleUpdate)
	kernelPlugin.AFCPlugin.RegisterModule(ConfigModuleType, ConfigModuleUpdate)
	kernelPlugin.AFCPlugin.RegisterModule(MapModuleType, MapModuleUpdate)
	kernelPlugin.AFCPlugin.RegisterModule(KernelModuleType, KernelModuleUpdate)
}

func (kernelPlugin *AFKernelPlugin) Uninstall() {

	kernelPlugin.AFCPlugin.DeregisterModule(MetaClassModuleName)
	kernelPlugin.AFCPlugin.DeregisterModule(ConfigModuleName)
	kernelPlugin.AFCPlugin.DeregisterModule(MapModuleName)
	kernelPlugin.AFCPlugin.DeregisterModule(KernelModuleName)
}

func (kernelPlugin *AFKernelPlugin) GetPluginName() string {
	return PluginName
}
