package kernelInterface

import (
	ark "github.com/ArkNX/ark-go/interface"
)

var AFIMetaClassModuleName string

type AFIMetaClassModule interface {
	ark.AFIModule
	Load() error
	AddClassCallBack(className string)
}
