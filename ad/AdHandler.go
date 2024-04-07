package ad

import (
	"dcard/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAd(c *gin.Context) {
	var adReqVo CreateAdReqVo
	err := c.BindJSON(&adReqVo)

	if err != nil {
		infra.Throw400IfError(c, err.Error())
		return
	}

	errorCode, errorMessage, adID := createAd(adReqVo)

	if errorCode == 0 {
		c.JSON(http.StatusCreated, gin.H{"AdID": adID})
	} else if errorCode == 400 {
		infra.Throw400IfError(c, errorMessage)
	} else {
		infra.Throw500IfError(c, errorMessage)
	}

}

func GetAd(c *gin.Context) {
	var vo GetAdReqVo
	if err := c.Bind(&vo); err != nil {
		infra.Throw400IfError(c, err.Error())
		return
	}

	errorCode, errorMessage, ads := getAd(vo)

	if errorCode == 0 {
		c.JSON(http.StatusOK, gin.H{"data": ads})
	} else if errorCode == 400 {
		infra.Throw400IfError(c, errorMessage)
	} else {
		infra.Throw500IfError(c, errorMessage)
	}
}
