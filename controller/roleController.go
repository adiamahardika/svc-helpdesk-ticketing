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

type roleController struct {
	roleService service.RoleServiceInterface
	logService  service.LogServiceInterface
}

func RoleController(roleService service.RoleServiceInterface, logService service.LogServiceInterface) *roleController {
	return &roleController{roleService, logService}
}

func (controller *roleController) GetRole(context *gin.Context) {

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	role, error := controller.roleService.GetRole()

	if error == nil {
		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":   status,
			"listRole": role,
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

	parse_status, _ := json.Marshal(status)
	parse_category, _ := json.Marshal(role)
	var result = fmt.Sprintf("{\"status\": %s, \"listRole\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *roleController) CreateRole(context *gin.Context) {

	var request model.CreateRoleRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var role []entity.Role

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

		role, error = controller.roleService.CreateRole(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":   status,
				"listRole": role,
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
	parse_category, _ := json.Marshal(role)
	var result = fmt.Sprintf("{\"status\": %s, \"listRole\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *roleController) UpdateRole(context *gin.Context) {

	var request model.UpdateRoleRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var role []model.GetRoleResponse

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

		role, error = controller.roleService.UpdateRole(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":   status,
				"listRole": role,
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
	parse_category, _ := json.Marshal(role)
	var result = fmt.Sprintf("{\"status\": %s, \"listRole\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
