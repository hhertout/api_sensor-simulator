package router

import (
	controller "api_sensor/controllers"
	"api_sensor/middleware"

	"github.com/gin-gonic/gin"
)

func authRoute(router *gin.Engine) {
	router.POST("api/auth", controller.CreateKey)
	router.GET("api/auth/validate", middleware.CheckCredentials, controller.GetAuth)
}
