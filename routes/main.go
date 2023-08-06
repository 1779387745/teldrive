package routes

import "github.com/gin-gonic/gin"

func GetRoutes(router *gin.Engine) {
	api := router.Group("/api")
	addFileRoutes(api)
}
