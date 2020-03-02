package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ArkNX/ark-go/util"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// 入口相关文件
const (
	// 导出so文件
	soEntryFile = `package main

import (
	ark "github.com/ArkNX/ark-go/interface"
	{{.PluginName}}Src "github.com/ArkNX/ark-go/plugin/{{.PluginName}}/src"
)

func main() {}

func DllEntryPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Register({{.PluginName}}Src.NewPlugin())
}

func DllExitPlugin(pPluginManager *ark.AFPluginManager) {
	pPluginManager.Deregister({{.PluginName}}Src.PluginName)
}`

	// 项目使用的插件入口函数
	pluginEntryFile = `package {{.PluginName}}

import (
	"github.com/ArkNX/ark-go/interface"
	{{.PluginName}}Src "github.com/ArkNX/ark-go/plugin/{{.PluginName}}/src"
)

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin({{.PluginName}}Src.PluginName, {{.PluginName}}Src.NewPlugin())
}`
)

// plugin about
const (
	pluginSrcFile = `package {{.PluginName}}Src

import (
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/util"
)

var PluginName = util.GetName((*AF{{.CPluginName}}Plugin)(nil))

type AF{{.CPluginName}}Plugin struct {
	ark.AFCPlugin
}

func NewPlugin() *AF{{.CPluginName}}Plugin {
	return &AF{{.CPluginName}}Plugin{AFCPlugin: ark.NewAFCPlugin()}
}

func ({{.PluginName}}Plugin *AF{{.CPluginName}}Plugin) Install() {
{{- range .Modules }}
	ark.RegisterModule({{.}}ModuleType, &{{$.PluginName}}Plugin.AFCPlugin, {{.}}ModuleUpdate)
{{- end }}
}

func ({{.PluginName}}Plugin *AF{{.CPluginName}}Plugin) Uninstall() {
{{ range .Modules }}
	ark.DeregisterModule({{.}}ModuleName, &{{$.PluginName}}Plugin.AFCPlugin)
{{- end }}
}

func ({{.PluginName}}Plugin *AF{{.CPluginName}}Plugin) GetPluginName() string {
	return PluginName
}`
)

// module about
const (
	moduleInterfaceFile = `package {{.PluginName}}Interface

import (
	ark "github.com/ArkNX/ark-go/interface"
)

var AFI{{.CModuleName}}ModuleName string

type AFI{{.CModuleName}}Module interface {
	ark.AFIModule
}`

	moduleSrcFile = `package {{.PluginName}}Src

import (
	"fmt"

	"github.com/ArkNX/ark-go/interface"
	{{.PluginName}}Interface "github.com/ArkNX/ark-go/plugin/{{.PluginName}}/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	{{.ModuleName}}ModuleType   = util.GetType((*AFC{{.CModuleName}}Module)(nil))
	{{.ModuleName}}ModuleName   = util.GetName((*AFC{{.CModuleName}}Module)(nil))
	{{.ModuleName}}ModuleUpdate = fmt.Sprintf("%p", (&AFC{{.CModuleName}}Module{}).Update) != fmt.Sprintf("%p", (&ark.AFCModule{}).Update)
)

func init() {
	{{.PluginName}}Interface.AFI{{.CModuleName}}ModuleName = util.GetName((*AFC{{.CModuleName}}Module)(nil))
}

type AFC{{.CModuleName}}Module struct {
	ark.AFCModule
	// other data
}

func ({{.ModuleName}}Module *AFC{{.CModuleName}}Module) Init() error {
	return nil
}`
)

var (
	plugin  string
	modules string
	out     string

	// dir about
	srcDir       string
	interfaceDir string
	pluginDir    string
)

type pluginTemplate struct {
	PluginName  string
	CPluginName string
	Modules     []string
}

type moduleTemplate struct {
	PluginName  string
	ModuleName  string
	CModuleName string
}

func main() {
	flag.StringVar(&plugin, "plugin", "", "name of the plugin to build")
	flag.StringVar(&modules, "modules", "", "module slice to build [ split by `,` ]")
	flag.StringVar(&out, "out", "", "out path of files")
	flag.Parse()

	var moduleArr []string
	for k, v := range []string{plugin, modules, out} {
		if len(v) == 0 && k != 1 {
			log.Fatalf("config `%s` is absents.", []string{"plugin", "modules", "out"}[k])
		}

		if k == 1 && len(v) != 0 {
			moduleArr = strings.Split(v, ",")
		}
	}

	// check the path is exist
	if !util.PathExists(out) {
		log.Fatalf("out path `%s` is invalid.", out)
	}

	out = filepath.Join(out, plugin)
	if !util.PathExists(out) {
		if err := os.Mkdir(out, os.ModePerm); err != nil {
			log.Fatalf("failed to mkdir : %s\n", out)
		}
	}

	// make dirs
	interfaceDir = filepath.Join(out, "interface")
	pluginDir = filepath.Join(out, "plugin")
	srcDir = filepath.Join(out, "src")
	for _, v := range []string{interfaceDir, pluginDir, srcDir} {
		if !util.PathExists(v) {
			if err := os.Mkdir(v, os.ModePerm); err != nil {
				log.Fatalf("failed to mkdir : %s\n", v)
			}
		}
	}

	genPlugin(pluginTemplate{
		PluginName:  plugin,
		CPluginName: Ucfirst(plugin),
		Modules:     moduleArr,
	})

	// --------------------module files--------------
	for _, module := range moduleArr {
		genModule(module, moduleTemplate{
			PluginName:  plugin,
			ModuleName:  module,
			CModuleName: Ucfirst(module),
		})
	}

}

func genPlugin(t pluginTemplate) {
	// gen files
	capitalPlugin := Ucfirst(plugin)
	pluginFileName := fmt.Sprintf("AF%sPlugin.go", capitalPlugin)

	// out/AFxxxPlugin.go
	pluginFile := filepath.Join(out, pluginFileName)
	if !util.PathExists(pluginFile) {
		str, err := parseTemplate(pluginEntryFile, t)
		if err != nil {
			log.Fatal(err)
		}

		if err := util.ForceWrite(pluginFile, []byte(str)); err != nil {
			log.Fatal(err)
		}
	}

	// out/src/AFxxxPlugin.go
	srcPluginFile := filepath.Join(srcDir, pluginFileName)
	if !util.PathExists(srcPluginFile) {
		str, err := parseTemplate(pluginSrcFile, t)
		if err != nil {
			log.Fatal(err)
		}

		if err := util.ForceWrite(srcPluginFile, []byte(str)); err != nil {
			log.Fatal(err)
		}
	}

	// out/plugin/AFxxxPlugin.go
	pluginPluginFile := filepath.Join(pluginDir, pluginFileName)
	if !util.PathExists(pluginPluginFile) {
		str, err := parseTemplate(soEntryFile, t)
		if err != nil {
			log.Fatal(err)
		}

		if err := util.ForceWrite(pluginPluginFile, []byte(str)); err != nil {
			log.Fatal(err)
		}
	}
}

func genModule(module string, mt moduleTemplate) {
	// gen files
	capitalModule := Ucfirst(module)
	moduleFileName := fmt.Sprintf("AF%sModule.go", capitalModule)

	// out/src/AFxxxModule.go
	srcModuleFile := filepath.Join(srcDir, moduleFileName)
	if !util.PathExists(srcModuleFile) {
		str, err := parseTemplate(moduleSrcFile, mt)
		if err != nil {
			log.Fatal(err)
		}

		if err := util.ForceWrite(srcModuleFile, []byte(str)); err != nil {
			log.Fatal(err)
		}
	}

	// out/interface/AFxxxModule.go
	interfaceModuleFile := filepath.Join(interfaceDir, moduleFileName)
	if !util.PathExists(interfaceModuleFile) {
		str, err := parseTemplate(moduleInterfaceFile, mt)
		if err != nil {
			log.Fatal(err)
		}

		if err := util.ForceWrite(interfaceModuleFile, []byte(str)); err != nil {
			log.Fatal(err)
		}
	}
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func parseTemplate(str string, data interface{}) (string, error) {
	t, err := template.New("test").Parse(str)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
