package config

import (
	//"context"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	ServerPort    configKey = "SERVER_PORT"
	MemoryMode    configKey = "MEMORY_MODE"
	TransportMode configKey = "TRANSPORT_MODE"

	DbName configKey = "DB_NAME"
	DbHost configKey = "DB_HOST"
	DbPort configKey = "DB_PORT"
	DbUser configKey = "DB_USER"

	ShortUrlPattern configKey = "SHORT_URL_PATTERN"

	LogLevel configKey = "LOG_LEVEL"
)

type (
	configKey   string
	configValue string

	environment struct {
		Env []configItem `yaml:"env"`
	}

	configItem struct {
		Name  configKey   `yaml:"name"`
		Value configValue `yaml:"value"`
	}
)

var configMap = make(map[configKey]configValue)

func init() {
	parseConfig()
}

func parseConfig() {
	conf, err := os.OpenFile("./config.yaml", os.O_RDONLY, os.ModeTemporary)

	if err != nil {
		panic(fmt.Sprintf("failed to open config file: %+v", err))
	}

	data, err := io.ReadAll(conf)
	if err != nil {
		panic(fmt.Sprintf("failed to read config file: %+v", err))
	}

	items := environment{}
	if err = yaml.Unmarshal(data, &items); err != nil {
		panic(fmt.Sprintf("failed to unmarshal from yaml: %+v", err))
	}

	for _, el := range items.Env {
		configMap[el.Name] = el.Value
	}
}

func Get(key configKey) configValue {
	v, ok := configMap[key]
	if !ok {
		panic(fmt.Sprintf("config key `%s` not found", key))
	}
	return v
}
