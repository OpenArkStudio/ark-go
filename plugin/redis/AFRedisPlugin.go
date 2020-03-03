package redis

import (
	"github.com/ArkNX/ark-go/interface"
	redisSrc "github.com/ArkNX/ark-go/plugin/redis/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(redisSrc.PluginName, redisSrc.NewPlugin())
}