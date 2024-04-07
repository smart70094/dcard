package infra

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Throw500IfError(c *gin.Context, message string) {
	log.Fatalln(message)
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Service abnormality, please contact the administrator"})
}

func Throw400IfError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}
