package options

import (
	"fmt"
	"os"
)

type Options struct {
	Port *string
	DbConnectionString *string
}

var ServiceSetting *Options = new(Options)

func InitSettings() *error {
	port, portExistsOk := os.LookupEnv("SettingsServicePort")
	dbConnectionString, dbConnectionStringOk := os.LookupEnv("SettingsServiceDbConnectionString")	

	if !portExistsOk || !dbConnectionStringOk {
		settings, err := readAppsettingsFile("appSettings.json")

		if err != nil {
			err := fmt.Errorf("Не удалось получить настройки приложения: " + err.Error())
			return &err
		}

		//TODO: читать файл настроек, если не найдены какие либо настройки 
		port = fmt.Sprintf("%v", settings["SettingsServicePort"])
		dbConnectionString = fmt.Sprintf("%v", settings["SettingsServiceDbConnectionString"])
	}

	ServiceSetting.Port = &port
	ServiceSetting.DbConnectionString = &dbConnectionString
	
	return nil
}