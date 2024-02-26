package main

import (
	"fmt"
	"gisogd/SettingsService/api"
	"gisogd/SettingsService/internal/options"
	"gisogd/SettingsService/internal/utils"
	"strings"

	"os"

	"go.uber.org/zap"
)

var Version string

//	@title						Settings service API
//	@version					1.0
//	@description				Сервис управления настройками информационной системы ГИСОГД
//	@BasePath					/v1
// 	@externalDocs.description  	OpenAPI
// 	@externalDocs.url          	https://swagger.io/resources/open-api/
func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			arg = strings.TrimLeft(arg, "-")
			switch arg {
			case "v", "version", "Version" : fmt.Println(Version)
			case "?", "help": fmt.Print(instructions) //TODO: write help instructions
			}
		}
		return
	}

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

const instructions string = `Rest API for managing microservices settings.

Usage: 

	SettingsService [arguments] 

Arguments:

	v, version, Version		returns current app version
	?, help, Help			shown this help
	
`
