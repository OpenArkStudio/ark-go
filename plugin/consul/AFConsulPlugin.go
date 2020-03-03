package consul

import (
	"github.com/ArkNX/ark-go/interface"
	consulSrc "github.com/ArkNX/ark-go/plugin/consul/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(consulSrc.PluginName, consulSrc.NewPlugin())
}