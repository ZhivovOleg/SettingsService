package dto

// @description запрос конкретной настройки
type GetConcreteOptionRequest struct {
	OptionPath string `json:"optionPath" binding:"required"`
}