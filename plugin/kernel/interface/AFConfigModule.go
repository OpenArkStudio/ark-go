package kernelInterface

import (
	ark "github.com/ArkNX/ark-go/interface"
)

var AFIConfigModuleName string

type AFIConfigModule interface {
	ark.AFIModule
}