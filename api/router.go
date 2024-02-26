package api

import (
	"gisogd/SettingsService/docs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitApi(port string, dbConnStr string) {

	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.New()
	
	//https://github.com/gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST,PATCH,PUT,GET,DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept-Encoding", "Authorization", "Cache-Control"},
		MaxAge:           time.Hour,
	}))
	
	settingsController := &Controller{}
	settingsController.InitController(router.Group("/v1"), "settings", dbConnStr)
	settingsController.initV1settingsController()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run("localhost:" + port)
}