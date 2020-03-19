package kernel

import (
	"github.com/ArkNX/ark-go/interface"
	kernelSrc "github.com/ArkNX/ark-go/plugin/kernel/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(kernelSrc.PluginName, kernelSrc.NewPlugin())
}