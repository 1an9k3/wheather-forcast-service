package service

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.POST("/health", health)
	route.GET("/health", health)
	route.POST("/predict")

}
