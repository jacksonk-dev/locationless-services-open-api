package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var services = []Service{
	{Title: "Vet", Description: "I treat animals"},
	{Title: "Dj", Description: "I spin decks"},
}

func GetServices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services)
}

func AddService(c *gin.Context) {
	var newService Service

	if err := c.BindJSON(&newService); err != nil {
		return
	}

	services = append(services, newService)
	c.IndentedJSON(http.StatusCreated, newService)
}
