package src

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule"
	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func init() {
	consulModule.ModuleName = ark.GetName((*AFCConsulModule)(nil))
	consulModule.ModuleType = ark.GetType((*AFCConsulModule)(nil))
	consulModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCConsulModule{}).Update).Pointer()).Name()
}

type AFCConsulModule struct {
	ark.AFCModule
	// other data
	config       *consulAPI.Config // consul config
	consulClient *consulAPI.Client // consul Client

	// service registration related
	registryConfig    *consulModule.ConsulServiceRegistryConfig
	consulCheckServer *http.Server
	consulCheckPort   int

	// service watcher related
	discoveryConfigs map[string]*consulModule.ConsulServiceDiscoveryConfig
}

func (consulModule_ *AFCConsulModule) Init() error {
	return nil
}

func (consulModule_ *AFCConsulModule) SetRegisterCenter(config *consulAPI.Config) error {
	client, err := consulAPI.NewClient(config)
	if err != nil {
		return err
	}

	consulModule_.config = config
	consulModule_.consulClient = client
	return nil
}

func (consulModule_ *AFCConsulModule) RegisterService(port int, c *consulModule.ConsulServiceRegistryConfig) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})
	consulModule_.registryConfig = c
	consulModule_.consulCheckPort = port
	consulModule_.consulCheckServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		if err := consulModule_.consulCheckServer.ListenAndServe(); err != nil {
			log.Printf("cnosul check server start fail : %v\n", err)
		}
	}()

	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = consulModule_.registryConfig.ID
	registration.Name = consulModule_.registryConfig.ServerType
	registration.Tags = consulModule_.registryConfig.Tags
	registration.Address = consulModule_.registryConfig.IP
	registration.Port = consulModule_.registryConfig.Port
	registration.Check = &consulAPI.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, consulModule_.consulCheckPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", // del this service in 15s after check fail
	}

	return consulModule_.consulClient.Agent().ServiceRegister(registration)
}

func (consulModule_ *AFCConsulModule) DeregisterService() error {
	if consulModule_.registryConfig == nil {
		return errors.New("service register config is absent")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	consulModule_.consulCheckServer.Shutdown(ctx)
	defer cancel()

	return consulModule_.consulClient.Agent().ServiceDeregister(consulModule_.registryConfig.ID)
}

func (consulModule_ *AFCConsulModule) GetHealthServices(c *consulModule.ConsulServiceDiscoveryConfig) (<-chan consulModule.ConsulAvailableServers, error) {
	noticeChan := make(chan consulModule.ConsulAvailableServers, 100)

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
			noticeChan <- consulModule.ConsulAvailableServers{
				ServerType: c.ServerType,
				Servers:    servers,
			}
		}
	}

	go func() {
		if err := plan.Run(consulModule_.config.Address); err != nil {
			log.Printf("Consul Watch Err: %+v\n", err)
		}
	}()

	return noticeChan, nil
}

func (consulModule_ *AFCConsulModule) GetKeyValue(key string) ([]byte, error) {
	pair, _, err := consulModule_.consulClient.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	if pair == nil {
		return nil, errors.New("kv pair is absent")
	}

	return pair.Value, nil
}

func (consulModule_ *AFCConsulModule) SetKeyValue(key string, value []byte) error {
	_, err := consulModule_.consulClient.KV().Put(&consulAPI.KVPair{
		Key:   key,
		Value: value,
	}, nil)
	return err
}

func (consulModule_ *AFCConsulModule) DelKeyValue(key string) error {
	_, err := consulModule_.consulClient.KV().Delete(key, nil)
	return err
}
