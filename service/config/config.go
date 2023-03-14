package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func MiddlewareConfig(c *gin.Context) {
	c.Set("config", InitConfig())
	c.Next()
}

func getCurrentAbsDir() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func InitConfig() *TomlConfig {
	var config *TomlConfig
	cfgPath, _ := filepath.Abs(path.Join(getCurrentAbsDir(), "../config.toml"))
	log.Printf("config path: %v", cfgPath)
	if _, err := toml.DecodeFile(cfgPath, &config); err != nil {
		log.Panicf("load config fail %v", err)
	}
	return config
}
