package main

import (
	"fmt"
	"strings"

	"github.com/ZhivovOleg/SettingsService/api"
	"github.com/ZhivovOleg/SettingsService/internal/options"
	"github.com/ZhivovOleg/SettingsService/internal/utils"

	"os"
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

	isDebug := false
	if env, _ := os.LookupEnv("SettingsServiceEnv"); env == "dev" {
		isDebug = true		
	}

	utils.InitializeLogger(isDebug)
	defer utils.Logger.Sync()
	
	initSettingsErr := options.InitSettings(isDebug)

	if initSettingsErr != nil {
		utils.Logger.Error("Can't init settings: " + initSettingsErr.Error())
		panic("Can't init settings: " + initSettingsErr.Error())
	}

	api.InitServer(*options.ServiceSetting.Port, *options.ServiceSetting.DBConnectionString, isDebug)
}

const instructions string = `Rest API for managing microservices settings.

Usage: 

	SettingsService [arguments] 

Arguments:

	v, version, Version		returns current app version
	?, help, Help			shown this help
	
`
