package main

import (
	"dcard/ad"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/ad", ad.CreateAd)
	r.GET("/ad", ad.GetAd)
	r.Run(":8080")
}
