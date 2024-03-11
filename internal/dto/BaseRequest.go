package dto

import "encoding/json"

// @description Базовый запрос
type BaseRequest struct {
	// Имя сервиса
	ServiceName string `json:"serviceName" binding:"required"`
	// Настройки в виде json-строки
	Options json.RawMessage `json:"options" binding:"required"`
}