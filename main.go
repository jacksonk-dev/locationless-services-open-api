package main

import (
	"petsAndAnimalsBE/packages/pets"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/api/pets", pets.GetPets)
	router.POST("/api/pets", pets.AddPet)

	router.Run(appPort)
}
