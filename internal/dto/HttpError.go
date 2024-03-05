package dto

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Ошибка парсинга"`
}