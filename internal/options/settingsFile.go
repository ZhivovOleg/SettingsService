package options

import (
	"encoding/json"
	"errors"
	"os"
)

type Options struct {
	Port *string				`json:"port"`
	DBConnectionString *string	`json:"dbConnectionString"`
}

func readAppsettingsFile(filename string) (*Options, error) {
	if _, err := os.Stat(filename); err == nil {
		jsonFile, err := os.Open(filename)		
		if err != nil {
			return nil, errors.New("Ошибка при открытии файла настроек приложения:" + err.Error())
		}		
		defer jsonFile.Close()

		result := &Options{}
		jsonParser := json.NewDecoder(jsonFile)
		jsonParser.Decode(&result)

		return result, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("не найден файл настроек приложения")
	}

	return nil, errors.New("ошибка при чтении файла настроек")
}