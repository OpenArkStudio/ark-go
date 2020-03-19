package module

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ArkNX/ark-go/tools/pluginBuilder/utils"
)

const str = `package src

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"{{.ProjectName}}/plugin/{{.PluginName}}Plugin/{{.ModuleName}}Module"
)

func init() {
	{{.ModuleName}}Module.ModuleName = ark.GetName((*AFC{{.UcfirstModuleName}}Module)(nil))
	{{.ModuleName}}Module.ModuleType = ark.GetType((*AFC{{.UcfirstModuleName}}Module)(nil))
	{{.ModuleName}}Module.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFC{{.UcfirstModuleName}}Module{}).Update).Pointer()).Name()
}

type AFC{{.UcfirstModuleName}}Module struct {
	ark.AFCModule
	// other data
}

func ({{.ModuleName}}Module *AFC{{.UcfirstModuleName}}Module) Init() error {
	return nil
}`

const iStr = `package {{.ModuleName}}Module

import (
	ark "github.com/ArkNX/ark-go/interface"
	"reflect"
)

var (
	ModuleName   string
	ModuleType   reflect.Type
	ModuleUpdate string
)

type AFI{{.UcfirstModuleName}}Module interface {
	ark.AFIModule
}`

type Config struct {
	ProjectName       string
	PluginName        string
	ModuleName        string
	UcfirstModuleName string
}

func BuildModule(c *Config, outPath string) error {
	srcStr, err := utils.ParseTemplate(str, c)
	if err != nil {
		return err
	}

	interfaceStr, err := utils.ParseTemplate(iStr, c)
	if err != nil {
		return err
	}

	// write to disk
	srcPath := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName), fmt.Sprintf("%sModule", c.ModuleName), "src")
	if !utils.PathExists(srcPath) {
		if err := os.MkdirAll(srcPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to mkdir : %s\n", srcPath)
		}
	}
	srcPath = filepath.Join(srcPath, fmt.Sprintf("AFC%sModule.go", c.UcfirstModuleName))
	if !utils.PathExists(srcPath) {
		if err := utils.Write(srcPath, []byte(srcStr)); err != nil {
			return err
		}
	} else {
		fmt.Printf("path %s is already exist.\n", srcPath)
		return nil
	}

	interfacePath := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName), fmt.Sprintf("%sModule", c.ModuleName), fmt.Sprintf("AFI%sModule.go", c.UcfirstModuleName))
	if utils.PathExists(interfacePath) {
		fmt.Printf("path %s is already exist.\n", interfacePath)
		return nil
	}

	if err := utils.Write(interfacePath, []byte(interfaceStr)); err != nil {
		return err
	}

	return nil
}
