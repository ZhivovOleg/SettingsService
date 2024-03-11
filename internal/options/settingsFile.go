package options

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

func readAppsettingsFile(fileName string) (string, string, error) {
	if _, err := os.Stat(fileName); err == nil {
		jsonFile, err := os.Open(fileName)		
		if err != nil {
			return "", "", errors.New("Ошибка при открытии файла настроек приложения:" + err.Error())
		}		
		defer jsonFile.Close()

		result := make(map[string]string)
		jsonParser := json.NewDecoder(jsonFile)
		jsonParser.Decode(&result)

		return result["settingsServicePort"], result["settingsServiceDbConnectionString"], nil
	} else if errors.Is(err, os.ErrNotExist) || errors.Is(err, &fs.PathError{}) {
		wd, _ := os.Getwd()
		return "", "", errors.New("не найден файл настроек приложения: " + wd + "/" + fileName)
	}

	return "", "", errors.New("ошибка при чтении файла настроек")
}