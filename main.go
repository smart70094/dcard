package main

import (
	"dcard/ad"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/v1/ad", ad.CreateAd)
	r.GET("/api/v1/ad", ad.GetAd)
	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
