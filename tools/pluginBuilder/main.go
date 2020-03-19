package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ArkNX/ark-go/tools/pluginBuilder/module"
	"github.com/ArkNX/ark-go/tools/pluginBuilder/plugin"
	"github.com/ArkNX/ark-go/tools/pluginBuilder/utils"
)

var (
	pluginName string
	modules    string
	out        string
	project    string
)

func main() {
	flag.StringVar(&pluginName, "plugin", "", "name of the plugin to build")
	flag.StringVar(&project, "project", "github.com/ArkNX/ark-go", "name of your project")
	flag.StringVar(&modules, "modules", "", "module slice to build [ split by `,` ]")
	flag.StringVar(&out, "out", "/Users/qinhan/go/src/ark-go/plugin", "out path of files")
	flag.Parse()

	// check args
	var moduleStr []string
	if len(modules) != 0 {
		moduleStr = strings.Split(modules, ",")
	}

	if len(project) == 0 {
		fmt.Println("project name is absent")
		return
	}

	if len(pluginName) == 0 {
		fmt.Println("plugin name is absent")
		return
	}

	if len(out) == 0 {
		fmt.Println("outPath is absent")
		return
	}

	// gen
	if err := plugin.BuildPlugin(&plugin.Config{
		ProjectName:       project,
		PluginName:        pluginName,
		ModuleNames:       moduleStr,
		UcfirstPluginName: utils.Ucfirst(pluginName),
	}, out); err != nil {
		log.Fatal(err)
	}

	for _, m := range moduleStr {
		if err := module.BuildModule(&module.Config{
			ProjectName:       project,
			PluginName:        pluginName,
			ModuleName:        m,
			UcfirstModuleName: utils.Ucfirst(m),
		}, out); err != nil {
			log.Fatal(err)
		}
	}
}
