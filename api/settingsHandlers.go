package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ZhivovOleg/SettingsService/internal/dal"
	"github.com/ZhivovOleg/SettingsService/internal/dto"
	"github.com/ZhivovOleg/SettingsService/internal/utils"

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
func (c *controller) getAllOptions(context *gin.Context) {
	requestContext := context.Request.Context()

	servicesSettings, err := dal.GetAllSettingsFromDB(&requestContext, &c.database)
	
	if err != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*err).Error())
		context.JSON(http.StatusInternalServerError, "Error on getting data from DB: " + (*err).Error())
		return
	}

	context.IndentedJSON(http.StatusOK, servicesSettings)
}

/* ----------------- SINGLE SERVICE OPERATIONS ------------------ */

//	@Summary		Get service settings
//	@Description	Get service settings JSON as string by service name
//	@Tags			settings
//	@Produce		json
//	@Param			serviceName		path		string	true	"Service name"
//	@Success		200	{object}	string
//	@Failure		400	{object}	dto.HttpError
//	@Failure		500	{object}	dto.HttpError
//	@Router			/settings/{serviceName} [get]
func (c *controller) getServiceOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()
	serviceSetting, err := dal.GetSettingsFromDB(&serviceName, &requestContext, &c.database)
	
	if err != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*err).Error())
		context.JSON(http.StatusInternalServerError, "Error on getting data from DB: " + (*err).Error())
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
func (c *controller) addServiceOptions(context *gin.Context) {
	requestContext := context.Request.Context()

	rawByteVal, bodyErr := io.ReadAll(context.Request.Body)

	if bodyErr != nil {
		context.JSON(http.StatusBadRequest, "Unmarshal error: " + bodyErr.Error())
	}
	
	tmpRawByteVal := make([]byte, 0)

	for i:=0; i<len(rawByteVal);i++ {
		tmpStr := string(rawByteVal[i])
		if tmpStr == "\n" || tmpStr == "\r" {
			continue
		}
		tmpRawByteVal = append(tmpRawByteVal, rawByteVal[i])
	}

	res := &dto.BaseRequest{}

	if len(res.ServiceName) <= 0 {
		context.JSON(http.StatusBadRequest, "Argument error: serviceName")
		return
	}

	jerr := json.Unmarshal(tmpRawByteVal, &res)

	if jerr != nil {
		context.JSON(http.StatusBadRequest, "Unmarshal error: " + jerr.Error())
		return
	}

	var resultOpt string
	serviceName := res.ServiceName
	options := string(res.Options)
	unqOptions, unquoteErr := strconv.Unquote(options)
	if unquoteErr == nil {
		resultOpt = unqOptions
	} else {
		resultOpt = options
	}

	if len(resultOpt) <= 0 {
		context.JSON(http.StatusBadRequest, "Argument error: options")
		return
	}

	insertErr := dal.InsertNewOptionsToDB(&serviceName, &resultOpt, &requestContext, &c.database)

	if insertErr != nil {
		context.JSON(http.StatusInternalServerError, "Error on set data to DB: " + (*insertErr).Error())
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
func (c *controller) deleteServiceWithOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	requestContext := context.Request.Context()

	err := dal.DeleteSettingsFromDB(&serviceName, &requestContext, &c.database)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Не удалось удалить настройки: " + (*err).Error())
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
func (c *controller) replaceServiceOptions(context *gin.Context) {
	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	var requestBody dto.ReplaceOptionsRequest
	err := json.NewDecoder(context.Request.Body).Decode(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Error on getting arguments from request body: " + err.Error())
		return
	}

	requestContext := context.Request.Context()
	updateErr := dal.ReplaceOptionsInDB(&serviceName, &requestBody.Options, &requestContext, &c.database)
	
	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, "Error on insert data to DB: " + (*updateErr).Error())
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
func (c *controller) updateSingleValue(context *gin.Context) {
	if mime := context.ContentType(); mime != "text/plain" {
		utils.Logger.Error("Request error: invalid MIME")
		context.JSON(http.StatusBadRequest, "Request error: invalid MIME")
		return
	}

	serviceName := strings.Trim(context.Param("serviceName"), "/")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	optionPath := strings.Trim(context.Param("path"), "/")
	if optionPath == "" {
		utils.Logger.Error("Argument error: option path not found")
		context.JSON(http.StatusBadRequest, "Argument error: option path not found")
		return
	}

	requestContext := context.Request.Context()

	byteVal, err := io.ReadAll(context.Request.Body)
	if err != nil {
		utils.Logger.Error("Error on getting value from request body: " + err.Error())
		context.JSON(http.StatusBadRequest, "Error on getting value from request body: " + err.Error())
		return
	}
	strVal := string(byteVal)

	dbErr := dal.UpdateOptionInDB(&serviceName, &optionPath, &strVal, &requestContext, &c.database)
	if dbErr != nil {
		utils.Logger.Error("Error while update data in DB: " + (*dbErr).Error())
		context.JSON(http.StatusBadRequest, "Error while update data in DB: " + (*dbErr).Error())
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
func (c *controller) getSingleValue(context *gin.Context) {
	serviceName := context.Param("serviceName")
	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	optionPath := context.Param("path")
	if optionPath == "" {
		utils.Logger.Error("Argument error: option path not found")
		context.JSON(http.StatusBadRequest, "Argument error: option path not found")
		return
	}

	requestContext := context.Request.Context()
	serviceSetting, serviceErr := dal.GetConcreteOptionFromDB(&serviceName, &optionPath, &requestContext, &c.database)
	
	if serviceErr != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*serviceErr).Error())
		context.JSON(http.StatusInternalServerError, "Error on getting data from DB: " + (*serviceErr).Error())
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
func (c *controller) deleteSingleValue(context *gin.Context) {
	serviceName := context.Param("serviceName")

	if serviceName == "" {
		utils.Logger.Error("Argument error: Service name not found")
		context.JSON(http.StatusBadRequest, "Argument error: Service name not found")
		return
	}

	optionPath := context.Param("path")

	if optionPath == "" {
		utils.Logger.Error("Argument error: option path not found")
		context.JSON(http.StatusBadRequest, "Argument error: option path not found")
		return
	}

	requestContext := context.Request.Context()
	serviceErr := dal.DeleteConcreteOptionFromDB(&serviceName, &optionPath, &requestContext, &c.database)
	
	if serviceErr != nil {
		utils.Logger.Error("Error on getting data from DB: " + (*serviceErr).Error())
		context.JSON(http.StatusInternalServerError, "Error on getting data from DB: " + (*serviceErr).Error())
		return
	}
	
	context.Status(http.StatusOK)
}