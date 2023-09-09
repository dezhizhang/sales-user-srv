package test

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

func Register(address string, id string, port int, name string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	registration := new(api.AgentServiceRegistration)
	registration.Address = address
	registration.ID = id
	registration.Port = port
	registration.Name = name
	registration.Tags = tags

	check := &api.AgentServiceCheck{
		HTTP:                           "http://127.0.0.1/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range services {
		fmt.Println(key)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	filter, err := client.Agent().ServicesWithFilter(`Service=="sales-user-srv"`)
	for key, _ := range filter {
		fmt.Println(key)
	}
}

func TestConsul(t *testing.T) {
	Register("127.0.0.1", "sales-user-srv", 8085, "sales-user-srv", []string{"sales-user-srv"})
	AllServices()
	FilterService()

}
