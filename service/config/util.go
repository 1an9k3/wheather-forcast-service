package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

func InitConfig() *TomlConfig {
	var config *TomlConfig
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		log.Panicf("load config fail")
	}
	return config
}
