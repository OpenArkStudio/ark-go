package main

import (
	ark "github.com/ArkNX/ark-go/interface"
	consulSrc "github.com/ArkNX/ark-go/plugin/consul/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(consulSrc.NewPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(consulSrc.PluginName)
}