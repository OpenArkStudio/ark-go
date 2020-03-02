package http

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/http/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(httpSrc.PluginName, httpSrc.NewPlugin())
}
