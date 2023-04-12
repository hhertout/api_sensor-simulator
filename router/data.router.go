package router

import (
	controller "api_sensor/controllers"

	"github.com/gin-gonic/gin"
)

func dataRoute(router *gin.Engine) {
	router.GET("api/air", controller.GetData)
}
