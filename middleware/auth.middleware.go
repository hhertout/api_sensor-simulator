package middleware

import (
	"api_sensor/config"
	"api_sensor/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCredentials(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var key schema.Auth

	result := config.DB.Where("`key`= ?", auth).First(&key)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
