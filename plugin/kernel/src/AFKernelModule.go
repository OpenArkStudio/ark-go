package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	KernelModuleType   = util.GetType((*AFCKernelModule)(nil))
	KernelModuleName   = util.GetName((*AFCKernelModule)(nil))
	KernelModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCKernelModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIKernelModuleName = util.GetName((*AFCKernelModule)(nil))
}

type AFCKernelModule struct {
	ark.AFCModule
	// other value
}

func (KernelModule *AFCKernelModule) Init() error {
	return nil
}
