package options

import (
	"fmt"
	"os"
)

var ServiceSetting *Options = new(Options)

// InitSettings - initialize setting from ENV or appSettings.json
func InitSettings() error {
	port, portExistsOk := os.LookupEnv("SettingsServicePort")
	dbConnectionString, dbConnectionStringOk := os.LookupEnv("SettingsServiceDbConnectionString")	

	if !portExistsOk || !dbConnectionStringOk {
		settings, err := readAppsettingsFile("configs/appSettings.json")

		if err != nil {
			err := fmt.Errorf("Не удалось получить настройки приложения: " + err.Error())
			return err
		}

		port = *settings.Port
		dbConnectionString = *settings.DBConnectionString
	}

	ServiceSetting.Port = &port
	ServiceSetting.DBConnectionString = &dbConnectionString
	
	return nil
}