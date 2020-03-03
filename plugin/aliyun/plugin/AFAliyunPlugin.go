package main

import (
	ark "github.com/ArkNX/ark-go/interface"
	aliyunSrc "github.com/ArkNX/ark-go/plugin/aliyun/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register(aliyunSrc.NewPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister(aliyunSrc.PluginName)
}