package main

import (
	"locationless-services/packages/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/services", services.GetServices)
	router.POST("/services", services.AddService)

	router.Run(appPort)
}
