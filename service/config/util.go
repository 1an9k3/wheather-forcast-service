package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"path"
	"path/filepath"
)

func InitConfig() *TomlConfig {
	var config *TomlConfig
	cfgPath, _ := filepath.Abs(path.Join(getCurrentAbsDir(), "../config.toml"))
	log.Printf("config path: %v", cfgPath)
	if _, err := toml.DecodeFile(cfgPath, &config); err != nil {
		log.Panicf("load config fail %v", err)
	}
	return config
}
