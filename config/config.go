package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Http Http
}

type Http struct {
	Servers  []Server
	Upstream Upstream
}

type Server struct {
	Port      uint32
	Name      string
	Locations []Location
}

type Location struct {
	Path     string
	Root     string
	Response Response
}

type Response struct {
	Status uint
	Data   string
}

type Upstream struct{}

func defaultConfig() Config {
	return Config{
		Http: Http{
			Servers:  []Server{},
			Upstream: Upstream{},
		},
	}
}

func Parse(path string) (Config, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return defaultConfig(), err
	}

	conf := Config{}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return defaultConfig(), err

	}

	fmt.Printf("Read config file:\n%v\n\n", conf)

	return conf, nil
}
