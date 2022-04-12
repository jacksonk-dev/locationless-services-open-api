package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	router.Run(appPort)
}
