package dto

// @description Замена настроек
type ReplaceOptionsRequest struct {
	// Настройки в виде json-строки
	Options string `json:"options" binding:"required" example:"{\"c\":\"ca\"}"`
}