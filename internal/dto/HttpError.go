package dto

// @description Результат при ошибке
type HTTPError struct {
	// Http-код ответа
	Code    int    `json:"code" example:"400"`
	// Текст ошибки
	Message string `json:"message" example:"Ошибка парсинга"`
}