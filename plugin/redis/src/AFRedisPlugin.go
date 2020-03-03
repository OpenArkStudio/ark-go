package redisSrc

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AFRedisPlugin)(nil))

type AFRedisPlugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AFRedisPlugin {
	return &AFRedisPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (redisPlugin *AFRedisPlugin) Install() {
	ark.RegisterModule(redisModuleType, &redisPlugin.AFCPlugin, redisModuleUpdate)
}

func (redisPlugin *AFRedisPlugin) Uninstall() {

	ark.DeregisterModule(redisModuleName, &redisPlugin.AFCPlugin)
}

func (redisPlugin *AFRedisPlugin) GetPluginName() string {
	return PluginName
}