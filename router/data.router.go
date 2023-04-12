package router

import (
	controller "api_sensor/controllers"
	"api_sensor/middleware"

	"github.com/gin-gonic/gin"
)

func dataRoute(router *gin.Engine) {
	router.GET("api/air", middleware.CheckCrendentials, controller.GetData)
	router.GET("api/air/latest", middleware.CheckCrendentials, controller.GetLatestData)
}
