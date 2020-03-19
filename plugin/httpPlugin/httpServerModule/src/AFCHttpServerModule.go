package src

import (
	"fmt"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func init() {
	httpServerModule.ModuleName = ark.GetName((*AFCHttpServerModule)(nil))
	httpServerModule.ModuleType = ark.GetType((*AFCHttpServerModule)(nil))
	httpServerModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCHttpServerModule{}).Update).Pointer()).Name()
}

type AFCHttpServerModule struct {
	ark.AFCModule
	// other data
	log logModule.AFILogModule
}

func (httpServerModule *AFCHttpServerModule) Init() error {
	m := httpServerModule.GetPluginManager().FindModule(logModule.ModuleName)
	logModule, ok := m.(logModule.AFILogModule)
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

//func (httpServerModule *AFCHttpServerModule) Update() error {
//	httpServerModule.log.GetLogger().WithField("test-key", "test-value").Warn("Update func in httpServerModule")
//	return nil
//}
