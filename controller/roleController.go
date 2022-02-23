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
			"status":    status,
			"list_role": role,
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
	var result = fmt.Sprintf("{\"status\": %s, \"list_role\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
