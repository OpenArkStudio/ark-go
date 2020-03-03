package consulSrc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"

	"github.com/ArkNX/ark-go/interface"
	consulInterface "github.com/ArkNX/ark-go/plugin/consul/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	consulModuleType   = util.GetType((*AFCConsulModule)(nil))
	consulModuleName   = util.GetName((*AFCConsulModule)(nil))
	consulModuleUpdate = fmt.Sprintf("%p", (&AFCConsulModule{}).Update) != fmt.Sprintf("%p", (&ark.AFCModule{}).Update)
)

func init() {
	consulInterface.AFIConsulModuleName = util.GetName((*AFCConsulModule)(nil))
}

type AFCConsulModule struct {
	ark.AFCModule
	// other data
	config       *consulAPI.Config // consul config
	consulClient *consulAPI.Client // consul Client

	// service registration related
	registryConfig    *consulInterface.ConsulServiceRegistryConfig
	consulCheckServer *http.Server
	consulCheckPort   int

	// service watcher related
	discoveryConfigs map[string]*consulInterface.ConsulServiceDiscoveryConfig
}

func (consulModule *AFCConsulModule) Init() error {
	return nil
}

func (consulModule *AFCConsulModule) SetRegisterCenter(config *consulAPI.Config) error {
	client, err := consulAPI.NewClient(config)
	if err != nil {
		return err
	}

	consulModule.config = config
	consulModule.consulClient = client
	return nil
}

func (consulModule *AFCConsulModule) RegisterService(port int, c *consulInterface.ConsulServiceRegistryConfig) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})
	consulModule.registryConfig = c
	consulModule.consulCheckPort = port
	consulModule.consulCheckServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		if err := consulModule.consulCheckServer.ListenAndServe(); err != nil {
			log.Printf("cnosul check server start fail : %v\n", err)
		}
	}()

	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = consulModule.registryConfig.ID
	registration.Name = consulModule.registryConfig.ServerType
	registration.Tags = consulModule.registryConfig.Tags
	registration.Address = consulModule.registryConfig.IP
	registration.Port = consulModule.registryConfig.Port
	registration.Check = &consulAPI.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, consulModule.consulCheckPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", // del this service in 15s after check fail
	}

	return consulModule.consulClient.Agent().ServiceRegister(registration)
}

func (consulModule *AFCConsulModule) DeregisterService() error {
	if consulModule.registryConfig == nil {
		return errors.New("service register config is absent")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	consulModule.consulCheckServer.Shutdown(ctx)
	defer cancel()

	return consulModule.consulClient.Agent().ServiceDeregister(consulModule.registryConfig.ID)
}

func (consulModule *AFCConsulModule) GetHealthServices(c *consulInterface.ConsulServiceDiscoveryConfig) (<-chan consulInterface.ConsulAvailableServers, error) {
	noticeChan := make(chan consulInterface.ConsulAvailableServers, 100)

	// build plan
	params := make(map[string]interface{})
	params["type"] = "service"
	params["service"] = c.ServerType
	params["tag"] = c.Tags
	plan, err := consulWatch.Parse(params)
	if err != nil {
		return nil, err
	}

	plan.Handler = func(index uint64, raw interface{}) {
		if raw == nil {
			return
		}

		if entries, ok := raw.([]*consulAPI.ServiceEntry); ok {
			var servers []string
			for _, entry := range entries {
				// healthy check fail, continue anyway
				if entry.Checks.AggregatedStatus() != consulAPI.HealthPassing {
					continue
				}
				servers = append(servers, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
			}
			noticeChan <- consulInterface.ConsulAvailableServers{
				ServerType: c.ServerType,
				Servers:    servers,
			}
		}
	}

	go func() {
		if err := plan.Run(consulModule.config.Address); err != nil {
			log.Printf("Consul Watch Err: %+v\n", err)
		}
	}()

	return noticeChan, nil
}

func (consulModule *AFCConsulModule) GetKeyValue(key string) ([]byte, error) {
	pair, _, err := consulModule.consulClient.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	if pair == nil {
		return nil, errors.New("kv pair is absent")
	}

	return pair.Value, nil
}

func (consulModule *AFCConsulModule) SetKeyValue(key string, value []byte) error {
	_, err := consulModule.consulClient.KV().Put(&consulAPI.KVPair{
		Key:   key,
		Value: value,
	}, nil)
	return err
}

func (consulModule *AFCConsulModule) DelKeyValue(key string) error {
	_, err := consulModule.consulClient.KV().Delete(key, nil)
	return err
}
