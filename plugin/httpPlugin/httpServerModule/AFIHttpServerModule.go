package httpServerModule

import (
	ark "github.com/ArkNX/ark-go/interface"
	"reflect"
)

var (
	ModuleName   string
	ModuleType   reflect.Type
	ModuleUpdate string
)

type AFIHttpServerModule interface {
	ark.AFIModule
	Start(port uint16) error
}
