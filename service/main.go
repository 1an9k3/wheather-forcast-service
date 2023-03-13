package service

import (
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.Use(ConfigMiddleware)
	route.POST("/health", health)
	route.GET("/health", health)
	route.POST("/predict", predict)

}
