package config

import (
	"encoding/json"
	"github.com/aidar-darmenov/message-delivery/interfaces"
	"log"
	"os"
)

type Configuration struct {
	HttpPort     int
	ListenerHost string
	ListenerPort int
	ListenerType string
}

//NewConfiguration read file, return configuration
func NewConfiguration(path string) interfaces.Configuration {
	var configuration Configuration
	configuration.InitConfigParams(path)
	return &configuration
}

//ReadFile Загрузка настроек из файла конфигураций
func (c *Configuration) InitConfigParams(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Configuration) Params() *Configuration {
	return c
}
