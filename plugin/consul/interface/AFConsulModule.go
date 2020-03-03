package consulInterface

import (
	consulAPI "github.com/hashicorp/consul/api"

	ark "github.com/ArkNX/ark-go/interface"
)

var AFIConsulModuleName string

// ConsulServiceRegistryConfig is service registry config
type ConsulServiceRegistryConfig struct {
	ID         string   // service id
	ServerType string   // service type
	IP         string   // service addr
	Port       int      // service port
	Tags       []string // service Tags
}

// ConsulServiceDiscoveryConfig is service watcher config
type ConsulServiceDiscoveryConfig struct {
	ServerType string   // target service type
	Tags       []string // target service tags
	Min        int      // minimum of available in wait
}

// ConsulAvailableServers defines available online services
type ConsulAvailableServers struct {
	ServerType string
	Servers    []string
}

type AFIConsulModule interface {
	ark.AFIModule
	SetRegisterCenter(config *consulAPI.Config) error
	RegisterService(port int, c *ConsulServiceRegistryConfig) error
	DeregisterService() error
	GetHealthServices(c *ConsulServiceDiscoveryConfig) (<-chan ConsulAvailableServers, error)
	GetKeyValue(key string) ([]byte, error)
	SetKeyValue(key string, value []byte) error
	DelKeyValue(key string) error
}
