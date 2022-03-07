package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	authService service.AuthServiceInterface
	logService  service.LogServiceInterface
}

func AuthController(authService service.AuthServiceInterface, logService service.LogServiceInterface) *authController {
	return &authController{authService, logService}
}

func (controller *authController) Login(context *gin.Context) {
	var request model.LoginRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var user model.LoginResponse

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

		user, error = controller.authService.Login(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":   status,
				"response": user,
				// "auth":     login,
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
	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	// parse_auth, _ := json.Marshal(login)
	var result = fmt.Sprintf("{\"status\": %s, \"response\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
