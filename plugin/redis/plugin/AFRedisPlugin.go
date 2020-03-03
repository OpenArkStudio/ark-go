package main

import (
	ark "github.com/ArkNX/ark-go/interface"
	redisSrc "github.com/ArkNX/ark-go/plugin/redis/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(redisSrc.NewPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(redisSrc.PluginName)
}