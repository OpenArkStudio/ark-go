package httpSrc

import (
	"fmt"
	"log"
	"net/http"

	"ark-go/interface"
	httpInterface "ark-go/plugin/http/interface"
	logInterface "ark-go/plugin/log/interface"
	"ark-go/util"
)

var (
	httpServerModuleType   = util.GetType((*AFCHttpServerModule)(nil))
	httpServerModuleName   = util.GetName((*AFCHttpServerModule)(nil))
	httpServerModuleUpdate = fmt.Sprintf("%p", (&AFCHttpServerModule{}).Update) != fmt.Sprintf("%p", (&ark.AFCModule{}).Update)
)

func init() {
	httpInterface.AFIHttpServerModuleName = httpServerModuleName
}

type AFCHttpServerModule struct {
	ark.AFCModule
	// other data
	log logInterface.AFILogModule
}

func (httpServerModule *AFCHttpServerModule) Init() error {
	m := httpServerModule.GetPluginManager().FindModule(logInterface.AFILogModuleName)
	logModule, ok := m.(logInterface.AFILogModule)
	if !ok {
		log.Fatal("failed to get log module in httpServer module")
	}
	httpServerModule.log = logModule
	return nil
}

func (httpServerModule *AFCHttpServerModule) PostInit() error {
	go httpServerModule.Start(9999)
	return nil
}

func (httpServerModule *AFCHttpServerModule) Start(port uint16) error {
	http.HandleFunc("/hello", HelloServer)
	httpServerModule.log.GetLogger().Warn("start http server on port : ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (httpServerModule *AFCHttpServerModule) Update() error {
	httpServerModule.log.GetLogger().WithField("test-key", "test-value").Warn("Update func in httpServerModule")
	return nil
}
