package dto

// @description Замена настроек
type ReplaceOptionsRequest struct {
	Options string `json:"options" binding:"required" example:"{\"c\":\"ca\"}"`
}