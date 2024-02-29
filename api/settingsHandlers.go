package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"gisogd/SettingsService/internal/dal"
	"gisogd/SettingsService/internal/dto"
	"gisogd/SettingsService/internal/utils"

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
func (c *Controller) GetAllOptions(context *gin.Context) {
	requestContext := context.Request.Context()

	servicesSettings, err := dal.GetAllSettingsFromDb(&requestContext, &c.database)
	
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
func (c *Controller) GetOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()
	serviceSetting, err := dal.GetSettingsFromDb(&serviceName, &requestContext, &c.database)
	
	if err != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*err).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*err).Error())
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
func (c *Controller) NewOption(context *gin.Context) {
	var requestBody dto.NewOptionsRequest
	err := json.NewDecoder(context.Request.Body).Decode(&requestBody)

	if err != nil {
		context.String(http.StatusBadRequest, "Error on getting arguments from request body: " + err.Error())
		return
	}

	requestContext := context.Request.Context()
	insertErr := dal.InsertNewOptionsToDb(&requestBody.ServiceName, &requestBody.Options, &requestContext, &c.database)

	if insertErr != nil {
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*insertErr).Error())
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
func (c *Controller) RemoveOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()

	err := dal.DeleteSettingsFromDb(&serviceName, &requestContext, &c.database)
	
	if err != nil {
		context.String(http.StatusInternalServerError, "Не удалось удалить настройки: " + (*err).Error())
		return
	}

	context.Status(http.StatusOK)
}

//	@Summary		Replace service settings
//	@Description	Completely replace service settings by service name
//	@Tags			settings
//	@Produce		application/json
//	@Param			serviceName		path	string						true	"Service name"
//	@Param			settings		body	dto.ReplaceOptionsRequest	true	"Service settings"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}  [put]
func (c *Controller) ReplaceOptions(context *gin.Context) {
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
	updateErr := dal.ReplaceOptionsInDb(&serviceName, &requestBody.Options, &requestContext, &c.database)
	
	if updateErr != nil {
		context.String(http.StatusInternalServerError, "Error on insert data to DB: " + (*updateErr).Error())
		return
	}

	context.Status(http.StatusOK)
}

/* ----------------- SINGLE FIELD OPERATIONS -------------------- */

//	@Summary		Update service settings
//	@Description	Update value for service settings by settings key. Set value in body with MIME text/plain
//	@Tags			settings
// 	@Accept			text/plain
//	@Produce		application/json
//	@Param			serviceName		path	string	true	"Service name"
//	@Param			path			path	string	true	"Option path"
//	@Param			value			body	string	true	"Option value"
//	@Success		200
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}/{path}  [patch]
func (c *Controller) UpdateOption(context *gin.Context) {
	if mime := context.ContentType(); mime != "text/plain" {
		utils.Logger.Error("Request error: invalid MIME")
		context.String(http.StatusBadRequest, "Request error: invalid MIME")
		return
	}

	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.String(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	optionPath := strings.Trim(context.Param("path"), "/")
	if optionPath == "" {
		utils.Logger.Error("Argument error: option path not found")
		context.String(http.StatusBadRequest, "Argument error: option path not found")
		return
	}

	requestContext := context.Request.Context()

	byteVal, err := io.ReadAll(context.Request.Body)
	if err != nil {
		utils.Logger.Error("Error on getting value from request body: " + err.Error())
		context.String(http.StatusBadRequest, "Error on getting value from request body: " + err.Error())
		return
	}
	strVal := string(byteVal)

	dbErr := dal.UpdateOptionInDb(&serviceName, &optionPath, &strVal, &requestContext, &c.database)
	if dbErr != nil {
		utils.Logger.Error("Error while update data in DB: " + (*dbErr).Error())
		context.String(http.StatusBadRequest, "Error while update data in DB: " + (*dbErr).Error())
		return
	}
	context.Status(http.StatusOK)
}

//	@Summary		Get concrete service option
//	@Description	Get service option as string by service name and option path
//	@Tags			settings
//	@Produce		application/json
//	@Param			serviceName		path		string	true	"Service name"
//	@Param			path			path		string	true	"Option path, comma-separated keys"
//	@Success		200	{object}	string
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}/{path} [get]
func (c *Controller) GetConcreteOption(context *gin.Context) {
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
	serviceSetting, serviceErr := dal.GetConcreteOptionFromDb(&serviceName, &optionPath, &requestContext, &c.database)
	
	if serviceErr != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*serviceErr).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*serviceErr).Error())
		return
	}
	
	context.JSON(http.StatusOK, serviceSetting)
}

//	@Summary		Delete concrete service option
//	@Description	Delete service option by service name and option path
//	@Tags			settings
//	@Produce		application/json
//	@Param			serviceName		path		string	true	"Service name"
//	@Param			path			path		string	true	"Option path, comma-separated keys"
//	@Success		200	{object}	string
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName}/{path} [delete]
func (c *Controller) DeleteConcreteOption(context *gin.Context) {
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
	serviceErr := dal.DeleteConcreteOptionFromDb(&serviceName, &optionPath, &requestContext, &c.database)
	
	if serviceErr != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*serviceErr).Error())
		context.String(http.StatusInternalServerError, "Error on getting data from DB: " + (*serviceErr).Error())
		return
	}
	
	context.Status(http.StatusOK)
}