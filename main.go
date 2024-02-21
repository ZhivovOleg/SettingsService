package main

import (
	"gisogd/SettingsService/api"
	"gisogd/SettingsService/dal"
	"gisogd/SettingsService/options"
	"gisogd/SettingsService/utils"
	"os"

	"go.uber.org/zap"
)

//	@title						Settings service API
//	@version					1.0
//	@description				Сервис управления настройками информационной системы ГИСОГД
//	@BasePath					/v1
// 	@externalDocs.description  	OpenAPI
// 	@externalDocs.url          	https://swagger.io/resources/open-api/
func main() {
	utils.InitializeLogger()
	if env, _ := os.LookupEnv("SettingsServiceEnv"); env != "dev" {
		utils.Logger, _ = zap.NewProduction()
	} else {
		utils.Logger, _ = zap.NewDevelopment()
	}
	defer utils.Logger.Sync()
	
	initSettingsErr := options.InitSettings()

	if initSettingsErr != nil {
		utils.Logger.Fatal("Can't init settings: " + (*initSettingsErr).Error())
	}
	
	pool, initPoolErr := dal.InitPool()
	if (*initPoolErr) != nil {
		utils.Logger.Fatal("Can't init database pool: " + (*initPoolErr).Error())
	}
	defer pool.Close()
	
	api.InitApi()
}