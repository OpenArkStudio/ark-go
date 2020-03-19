package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	MapModuleType   = util.GetType((*AFCMapModule)(nil))
	MapModuleName   = util.GetName((*AFCMapModule)(nil))
	MapModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCMapModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIMapModuleName = util.GetName((*AFCMapModule)(nil))
}

type AFCMapModule struct {
	ark.AFCModule
	// other value
}

func (MapModule *AFCMapModule) Init() error {
	return nil
}
