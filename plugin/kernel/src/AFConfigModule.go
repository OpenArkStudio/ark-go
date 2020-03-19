package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	ConfigModuleType   = util.GetType((*AFCConfigModule)(nil))
	ConfigModuleName   = util.GetName((*AFCConfigModule)(nil))
	ConfigModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCConfigModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIConfigModuleName = util.GetName((*AFCConfigModule)(nil))
}

type AFCConfigModule struct {
	ark.AFCModule
	// other value
}

func (ConfigModule *AFCConfigModule) Init() error {
	return nil
}
