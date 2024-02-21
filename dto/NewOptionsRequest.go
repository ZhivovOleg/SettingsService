package dto

// @description Создать значение для сервиса
type NewOptionsRequest struct {
	// Имя сервиса
	ServiceName string `json:"serviceName" binding:"required"`
	// Настройки в виде json-строки 
	Options string `json:"options" binding:"required" example:"{\"a\":\"a3\"}"`
}