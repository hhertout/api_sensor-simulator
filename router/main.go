package router

import (
	"github.com/gin-gonic/gin"
)

func Main() {
	router := gin.New()

	authRoute(router)
	dataRoute(router)

	router.Run()
}
