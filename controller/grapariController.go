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

type grapariController struct {
	grapariService service.GrapariServiceInterface
	logService     service.LogServiceInterface
}

func GrapariController(grapariService service.GrapariServiceInterface, logService service.LogServiceInterface) *grapariController {
	return &grapariController{grapariService, logService}
}

func (controller *grapariController) GetGrapari(context *gin.Context) {

	var request *model.GetGrapariRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var grapari []entity.MsGrapari

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = &model.StandardResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		grapari, error = controller.grapariService.GetGrapari(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  status,
				"content": grapari,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			status = &model.StandardResponse{
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
	parse_grapari, _ := json.Marshal(grapari)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s", string(parse_status), string(parse_grapari))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
