package main

import (
	"ark-go/interface"
	"ark-go/plugin/log/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(logSrc.NewAFLogPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(logSrc.LogPluginName)
}
