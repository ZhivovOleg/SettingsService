package options

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func readAppsettingsFile(filename string) (map[string]interface{}, error) {
	if _, err := os.Stat(filename); err == nil {
		jsonFile, err := os.Open(filename)
		
		if err != nil {
			return nil, errors.New("Ошибка при открытии файла настроек приложения:" + err.Error())
		}
		
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		return result, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("не найден файл настроек приложения")
	}

	return nil, errors.New("ошибка при чтении файла настроек")
}