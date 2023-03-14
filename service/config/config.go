package config

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareConfig(c *gin.Context) {
	c.Set("config", InitConfig())
	c.Next()
}
