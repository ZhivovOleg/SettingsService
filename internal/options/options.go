package options

import (
	"fmt"
	"os"
)

type Options struct {
	Port *string				`json:"settingsServicePort"`
	DBConnectionString *string	`json:"settingsServiceDbConnectionString"`
}

var ServiceSetting *Options = new(Options)

// InitSettings - initialize setting from ENV or appSettings.json
func InitSettings(isDebug bool) error {
	port, portExistsOk := os.LookupEnv("SettingsServicePort")
	dbConnectionString, dbConnectionStringOk := os.LookupEnv("SettingsServiceDbConnectionString")	

	if !portExistsOk || !dbConnectionStringOk {
		var err error
		if isDebug {
			port, dbConnectionString, err = readAppsettingsFile("../../configs/appSettings.json")
		} else {
			port, dbConnectionString, err = readAppsettingsFile("appSettings.json")
		}

		if err != nil {
			err := fmt.Errorf("Не удалось получить настройки приложения: " + err.Error())
			return err
		}
	}

	ServiceSetting.Port = &port
	ServiceSetting.DBConnectionString = &dbConnectionString
	
	return nil
}