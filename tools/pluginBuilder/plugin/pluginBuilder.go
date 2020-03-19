package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ArkNX/ark-go/tools/pluginBuilder/utils"
)

const str = `package {{.PluginName}}Plugin

import (
	ark "github.com/ArkNX/ark-go/interface"
{{- range .ModuleNames }}
	"{{$.ProjectName}}/plugin/{{$.PluginName}}Plugin/{{.}}Module"
{{- end }}
{{ range .ModuleNames }}
	_ "{{$.ProjectName}}/plugin/{{$.PluginName}}Plugin/{{.}}Module/src"
{{- end }}
)

var PluginName = ark.GetName((*AF{{.UcfirstPluginName}}Plugin)(nil))

type AF{{.UcfirstPluginName}}Plugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AF{{.UcfirstPluginName}}Plugin {
	return &AF{{.UcfirstPluginName}}Plugin{AFCPlugin: ark.NewAFCPlugin()}
}

func ({{.PluginName}}Plugin *AF{{.UcfirstPluginName}}Plugin) Install() {
{{- range .ModuleNames }}
	{{$.PluginName}}Plugin.AFCPlugin.RegisterModule({{.}}Module.ModuleType, {{.}}Module.ModuleUpdate)
{{- end }}
}

func ({{.PluginName}}Plugin *AF{{.UcfirstPluginName}}Plugin) Uninstall() {
{{- range .ModuleNames }}
	{{$.PluginName}}Plugin.AFCPlugin.DeregisterModule({{.}}Module.ModuleName)
{{- end }}
}

func ({{.PluginName}}Plugin *AF{{.UcfirstPluginName}}Plugin) GetPluginName() string {
	return PluginName
}`

type Config struct {
	ProjectName       string
	PluginName        string
	UcfirstPluginName string
	ModuleNames       []string
}

func BuildPlugin(c *Config, outPath string) error {
	pluginStr, err := utils.ParseTemplate(str, c)
	if err != nil {
		return err
	}

	path := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName))
	if !utils.PathExists(path) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to mkdir : %s\n", path)
		}
	}

	path = filepath.Join(path, fmt.Sprintf("AF%sPlugin.go", c.UcfirstPluginName))
	if utils.PathExists(path) {
		fmt.Printf("path %s is already exist.\n", path)
		return nil
	}

	if err := utils.Write(path, []byte(pluginStr)); err != nil {
		return err
	}
	return nil
}
