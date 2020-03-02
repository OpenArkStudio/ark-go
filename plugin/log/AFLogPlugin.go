package log

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/log/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(logSrc.LogPluginName, logSrc.NewAFLogPlugin())
}

// if you want to use `so`
//func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
//	pPluginManager.Register(logSrc.NewAFLogPlugin())
//}
//
//func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
//	pPluginManager.Deregister(logSrc.LogPluginName)
//}
