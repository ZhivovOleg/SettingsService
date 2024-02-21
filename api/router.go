package api

import (
	"fmt"
	"gisogd/SettingsService/docs"
	"gisogd/SettingsService/options"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitApi() {

	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.New()
	
	//https://github.com/gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST,PATCH,GET,DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		MaxAge:           12 * time.Hour,
	}))

	v1settingsController(router.Group("/v1"))
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run(fmt.Sprintf("localhost:%s", (*options.ServiceSetting.Port)))
}