package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	MetaClassModuleType   = util.GetType((*AFCMetaClassModule)(nil))
	MetaClassModuleName   = util.GetName((*AFCMetaClassModule)(nil))
	MetaClassModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCMetaClassModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIMetaClassModuleName = util.GetName((*AFCMetaClassModule)(nil))
}

type AFCMetaClassModule struct {
	ark.AFCModule
	// other value
}

func (MetaClassModule *AFCMetaClassModule) Init() error {
	return nil
}
