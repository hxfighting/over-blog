package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	Config = getConfig()
)

func getConfig() *toml.Tree {
	config, err := toml.LoadFile("./config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}
	return config
}

func GetConfig(key string) interface{} {
	return Config.Get(key)
}
