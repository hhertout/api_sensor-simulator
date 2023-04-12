package controller

import (
	"api_sensor/config"
	"api_sensor/core/tools"
	"api_sensor/schema"

	"github.com/gin-gonic/gin"
	"github.com/mintance/go-uniqid"
)

func GetAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Authentification successful",
	})
}

func CreateKey(c *gin.Context) {
	now := tools.GetStringTimestamp()
	id := uniqid.New(uniqid.Params{now, true})

	config.DB.Create(&schema.Auth{
		Key: id,
	})

	c.JSON(200, gin.H{
		"credentials": id,
	})
}
