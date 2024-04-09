package main

import (
	"dcard/ad"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1", hello)
	r.POST("/api/v1/ad", ad.CreateAd)
	r.GET("/api/v1/ad", ad.GetAd)
	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}
