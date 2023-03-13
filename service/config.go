package service

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
)

func ConfigMiddleware(c *gin.Context) {
	c.Set("config", initConfig())
	c.Next()
}

func initConfig() *TomlConfig {
	var config *TomlConfig
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		log.Panicf("load config fail")
	}
	return config
}
