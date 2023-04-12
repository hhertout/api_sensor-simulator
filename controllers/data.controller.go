package controller

import (
	"api_sensor/config"
	"api_sensor/schema"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	var data []schema.AirData

	result := config.DB.Preload("Data").Limit(60).Find(&data)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})

		return
	}

	c.JSON(200, data)
}

func GetLatestData(c *gin.Context) {
	var data schema.AirData

	result := config.DB.Preload("Data").First(&data)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}
