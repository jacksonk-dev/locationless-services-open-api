package pets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
