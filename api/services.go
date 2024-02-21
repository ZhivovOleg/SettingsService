package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"gisogd/SettingsService/dal"
	"gisogd/SettingsService/dto"
	"gisogd/SettingsService/utils"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get all service settings
//	@Description	Get all service settings JSON as map
//	@Tags			settings
//	@Produce		json
//	@Success		200	{object}	array
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings [get]
func GetAllOptions(context *gin.Context) {
	requestContext := context.Request.Context()

	servicesSettings, err := dal.GetAllSettingsFromDb(&requestContext)
	
	if err != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*err).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*err).Error())
		return
	}

	context.IndentedJSON(http.StatusOK, servicesSettings)
}

//	@Summary		Get service settings
//	@Description	Get service settings JSON as string by service name
//	@Tags			settings
//	@Produce		json
//	@Param			serviceName		path		string	true	"Service name"
//	@Success		200	{object}	string
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName} [get]
func GetOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()
	serviceSetting, err := dal.GetSettingsFromDb(&serviceName, &requestContext)
	
	if err != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*err).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*err).Error())
		return
	}

	context.JSON(http.StatusOK, serviceSetting)
}

//	@Summary		Get concrete service option
//	@Description	Get service option as string by service name and  option path
//	@Tags			settings
//  @Accept       	json
//	@Produce		json
//	@Param			serviceName		path		string	true	"Service name"
//	@Param			path			path		string	true	"Option path, comma-separated keys"
//	@Success		200	{object}	string
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}/{path} [get]
func GetConcreteOption(context *gin.Context) {
	serviceName := context.Param("serviceName")

	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	optionPath := context.Param("path")

	if optionPath == "" {
		utils.Logger.Error("Argument error: option path not found")
		context.String(http.StatusBadRequest, "Argument error: option path not found")
		return
	}

	requestContext := context.Request.Context()
	serviceSetting, serviceErr := dal.GetConcreteOptionFromDb(&serviceName, &optionPath, &requestContext)
	
	if serviceErr != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*serviceErr).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*serviceErr).Error())
		return
	}
	
	context.JSON(http.StatusOK, serviceSetting)
}

//	@Summary		Add complete settings
//	@Description	Add complete settings for new service as string by service name
//	@Tags			settings
//  @Accept       	json
//	@Produce		json
//	@Param			options		body	dto.NewOptionsRequest	true	"Options DTO"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings [post]
func NewOption(context *gin.Context) {
	var requestBody dto.NewOptionsRequest
	err := json.NewDecoder(context.Request.Body).Decode(&requestBody)

	if err != nil {
		context.String(http.StatusBadRequest, "Error on getting arguments from request body: " + err.Error())
		return
	}

	requestContext := context.Request.Context()
	insertErr := dal.InsertNewOptionsToDb(&requestBody.ServiceName, &requestBody.Options, &requestContext)

	if insertErr != nil {
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + err.Error())
		return
	}
	
	context.Status(http.StatusOK)
}

//	@Summary		Complete remove service settings
//	@Description	Complete remove service settings by service name
//	@Tags			settings
//	@Produce		json
//	@Param			serviceName		path	string	true	"Service name"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}  [delete]
func RemoveOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()

	err := dal.DeleteSettingsFromDb(&serviceName, &requestContext)
	
	if err != nil {
		context.String(http.StatusInternalServerError, "Не удалось удалить настройки: " + (*err).Error())
		return
	}

	context.Status(http.StatusOK)
}

//	@Summary		Replace service settings
//	@Description	Completely replace service settings by service name
//	@Tags			settings
//	@Param			serviceName		path	string						true	"Service name"
//	@Param			settings		body	dto.ReplaceOptionsRequest	true	"Service settings"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}  [put]
func ReplaceOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	var requestBody dto.ReplaceOptionsRequest
	err := json.NewDecoder(context.Request.Body).Decode(&requestBody)

	if err != nil {
		context.String(http.StatusBadRequest, "Error on getting arguments from request body: " + err.Error())
		return
	}

	requestContext := context.Request.Context()
	updateErr := dal.ReplaceOptionsInDb(&serviceName, &requestBody.Options, &requestContext)
	
	if updateErr != nil {
		context.String(http.StatusInternalServerError, "Error on insert data to DB: " + err.Error())
		return
	}

	context.Status(http.StatusOK)
}

//	@Summary		Update service settings
//	@Description	Update value for service settings by settings key
//	@Tags			settings
//	@Param			serviceName		path	string						true	"Service name"
//	@Param			update			body	dto.UpdateOptionRequest		true	"Service settings"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}  [patch]
func UpdateOption(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()

	var requestBody dto.UpdateOptionRequest
	err := json.NewDecoder(context.Request.Body).Decode(&requestBody)
	if err != nil {
		utils.Logger.Error("Error on getting arguments from request body: " + err.Error())
		context.String(http.StatusBadRequest, "Error on getting arguments from request body: " + err.Error())
		return
	}

	err = dal.UpdateOptionInDb(&serviceName, &requestBody.OptionPath, &requestBody.OptionValue, &requestContext)
	if err != nil {
		utils.Logger.Error("Error while update data in DB: " + err.Error())
		context.String(http.StatusBadRequest, "Error while update data in DB: " + err.Error())
		return
	}
	context.Status(http.StatusOK)
}