package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type emailNotifController struct {
	emailNotifService service.EmailNotifServiceInterface
	logService        service.LogServiceInterface
}

func EmailNotifController(emailNotifService service.EmailNotifServiceInterface, logService service.LogServiceInterface) *emailNotifController {
	return &emailNotifController{emailNotifService, logService}
}

func (controller *emailNotifController) CreateEmailNotif(context *gin.Context) {

	var email_notif entity.EmailNotif

	error := context.ShouldBind(&email_notif)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		email_notif, error = controller.emailNotifService.CreateEmailNotif(email_notif)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":      status,
				"email_notif": email_notif,
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
	}
	parse_request, _ := json.Marshal(email_notif)
	parse_status, _ := json.Marshal(status)
	parse_email_notif, _ := json.Marshal(email_notif)
	var result = fmt.Sprintf("{\"status\": %s, \"email_notif\": %s}", string(parse_status), string(parse_email_notif))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
