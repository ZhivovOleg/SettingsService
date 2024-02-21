package dto

// @description Изменить значение одного поля. Создает поле, если его не существует
type UpdateOptionRequest struct {
	// Новое значение
	OptionValue string `json:"optionValue" binding:"required"`
	// Поле для изменения
	OptionPath string `json:"optionPath" binding:"required" example:"a/b/c/1"`
}