package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

var (
	Config     *toml.Tree
	ConfigPath = ""
)

func NewConfig() {
	config, err := toml.LoadFile(ConfigPath + "/config.toml")
	if err != nil {
		log.Fatalln("缺少配置文件，", err.Error())
	}
	Config = config
}

func GetConfig(key string) interface{} {
	return Config.Get(key)
}
