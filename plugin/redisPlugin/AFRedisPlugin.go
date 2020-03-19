package redisPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/redisPlugin/redisModule"

	_ "github.com/ArkNX/ark-go/plugin/redisPlugin/redisModule/src"
)

var PluginName = ark.GetName((*AFRedisPlugin)(nil))

type AFRedisPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFRedisPlugin {
	return &AFRedisPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (redisPlugin *AFRedisPlugin) Install() {
	redisPlugin.AFCPlugin.RegisterModule(redisModule.ModuleType, redisModule.ModuleUpdate)
}

func (redisPlugin *AFRedisPlugin) Uninstall() {
	redisPlugin.AFCPlugin.DeregisterModule(redisModule.ModuleName)
}

func (redisPlugin *AFRedisPlugin) GetPluginName() string {
	return PluginName
}
