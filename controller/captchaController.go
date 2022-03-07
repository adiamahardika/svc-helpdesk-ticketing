package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type captchaController struct {
	captchaService service.CaptchaServiceInterface
	logService     service.LogServiceInterface
}

func CaptchaController(captchaService service.CaptchaServiceInterface, logService service.LogServiceInterface) *captchaController {
	return &captchaController{captchaService, logService}
}

var store = base64Captcha.DefaultMemStore

func (controller *captchaController) GenerateCaptcha(context *gin.Context) {

	decoder := json.NewDecoder(context.Request.Body)
	var request model.ConfigJsonBody
	error := decoder.Decode(&request)
	if error != nil {
		log.Println(error)
	}
	defer context.Request.Body.Close()

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var response map[string]interface{}
	var id string
	var b64s string

	id, b64s, error = controller.captchaService.GenerateCaptcha(request)

	response = map[string]interface{}{"image": b64s, "captchaId": id}

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":   status,
			"response": response,
		})

	} else {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})

	}

	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	var result = fmt.Sprintf("{\"status\": %s, \"response\": %s}", string(parse_status), response)
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
