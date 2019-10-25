package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

var (
	Config = getConfig()
)

func getConfig() *toml.Tree {
	dir, _ := os.Executable()
	config, err := toml.LoadFile(dir + "/config.toml")
	if err != nil {
		fmt.Println("缺少配置文件，", err.Error())
	}
	return config
}

func GetConfig(key string) interface{} {
	return Config.Get(key)
}
