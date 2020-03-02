package main

import (
	"ark-go/interface"
	"ark-go/plugin/http/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(httpSrc.NewAFHttpPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(httpSrc.HttpPluginName)
}
