package kernelInterface

import (
	ark "github.com/ArkNX/ark-go/interface"
)

var AFIKernelModuleName string

type AFIKernelModule interface {
	ark.AFIModule
}