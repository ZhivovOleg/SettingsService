package main

import (
	"gisogd/SettingsService/api"
	"gisogd/SettingsService/internal/options"
	"gisogd/SettingsService/internal/utils"

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
		utils.Logger.Error("Can't init settings: " + (*initSettingsErr).Error())
		panic("Can't init settings: " + (*initSettingsErr).Error())
	}

	api.InitApi(*options.ServiceSetting.Port, *options.ServiceSetting.DbConnectionString)
}
