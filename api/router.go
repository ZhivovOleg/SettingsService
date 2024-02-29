package api

import (
	"gisogd/SettingsService/docs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitApi(port string, dbConnStr string) {

	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.New()
	
	//init cors
	//https://github.com/gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST,PATCH,PUT,GET,DELETE"},
		AllowHeaders:     []string{"Content-Type"}, //"Accept-Encoding", "Authorization", "Cache-Control"},
		MaxAge:           time.Hour,
	}))
	
	//init front
	router.Use(static.Serve("/", static.LocalFile("web", false)))
	
	//init swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//init api
	api := router.Group("/api")

	//init version
	v1 := api.Group("/v1")

	//init settings controller
	settingsController := &Controller{}
	settingsController.InitController(v1, "settings", dbConnStr)
	settingsController.initV1settingsController()

    router.Run("localhost:" + port)
}