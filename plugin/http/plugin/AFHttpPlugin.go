package main

import (
	ark "github.com/ArkNX/ark-go/interface"
	httpSrc "github.com/ArkNX/ark-go/plugin/http/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(httpSrc.NewPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(httpSrc.PluginName)
}
