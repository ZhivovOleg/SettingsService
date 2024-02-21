package api

import "github.com/gin-gonic/gin"

func v1settingsController(v1router *gin.RouterGroup) {
	settings := v1router.Group("/settings")
	settings.GET("/", GetAllOptions)
	settings.POST("/", NewOption)
    settings.GET("/:serviceName", GetOptions)
    settings.GET("/:serviceName/:path", GetConcreteOption)
	settings.DELETE("/:serviceName", RemoveOptions)
	settings.PUT("/:serviceName", ReplaceOptions)
	settings.PATCH("/:serviceName", UpdateOption)
}