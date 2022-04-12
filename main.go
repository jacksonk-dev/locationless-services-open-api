package main

import (
	"petsAndAnimalsBE/packages/pets"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/api", pets.Get)

	router.Run(appPort)
}
