package log

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/log/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(logSrc.PluginName, logSrc.NewPlugin())
}
