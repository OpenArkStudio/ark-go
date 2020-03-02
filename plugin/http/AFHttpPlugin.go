package http

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/http/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(httpSrc.HttpPluginName, httpSrc.NewAFHttpPlugin())
}

// if you want to use `so`
//func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
//	pPluginManager.Register(httpSrc.NewAFHttpPlugin())
//}
//
//func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
//	pPluginManager.Deregister(httpSrc.HttpPluginName)
//}
