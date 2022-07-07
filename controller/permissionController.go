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

type permissionController struct {
	permissionService service.PermissionServiceInterface
	logService        service.LogServiceInterface
}

func PermissionController(permissionService service.PermissionServiceInterface, logService service.LogServiceInterface) *permissionController {
	return &permissionController{permissionService, logService}
}

func (controller *permissionController) GetPermission(context *gin.Context) {
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	permission, error := controller.permissionService.GetPermission()

	if error == nil {
		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":         status,
			"listPermission": permission,
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

	parse_status, _ := json.Marshal(status)
	parse_permission, _ := json.Marshal(permission)
	var result = fmt.Sprintf("{\"status\": %s, \"listPermission\": %s}", string(parse_status), string(parse_permission))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
