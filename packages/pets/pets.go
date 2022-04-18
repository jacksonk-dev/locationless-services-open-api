package pets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pet struct {
	PetType     string `json:"type"`
	Description string `json:"description"`
}

var pets = []Pet{
	{PetType: "dog", Description: "The best dog"},
	{PetType: "cat", Description: "A fuckin' cat"},
}

func GetPets(ginContext *gin.Context) {
	ginContext.IndentedJSON(http.StatusOK, pets)
}

func AddPet(ginContext *gin.Context) {
	var newPet Pet

	if err := ginContext.BindJSON(&newPet); err != nil {
		return
	}

	pets = append(pets, newPet)
	ginContext.IndentedJSON(http.StatusCreated, newPet)
}
