package httpInterface

import "github.com/ArkNX/ark-go/interface"

var AFIHttpServerModuleName string

type AFIHttpServerModule interface {
	ark.AFIModule
	Start(port uint16) error
}
