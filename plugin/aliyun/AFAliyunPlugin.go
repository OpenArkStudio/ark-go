package aliyun

import (
	"github.com/ArkNX/ark-go/interface"
	aliyunSrc "github.com/ArkNX/ark-go/plugin/aliyun/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(aliyunSrc.PluginName, aliyunSrc.NewPlugin())
}