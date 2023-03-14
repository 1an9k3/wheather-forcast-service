package main

import (
	"github.com/fyk1234/wheather-forcast-service/service/config"
	"github.com/fyk1234/wheather-forcast-service/service/processor"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	tomlCfg := config.InitConfig()
	route := gin.Default()
	route.Use(config.MiddlewareConfig)
	route.POST("/health", processor.Health)
	route.GET("/health", processor.Health)
	route.POST("/predict", processor.Predict)
	err := route.Run(tomlCfg.Addr)
	if err != nil {
		log.Fatalf("error start server: %v", err)
	}
}
