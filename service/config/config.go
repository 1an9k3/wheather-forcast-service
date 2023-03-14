package config

import (
	"github.com/gin-gonic/gin"
	"path"
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
