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

type subCategoryController struct {
	subCategoryService service.SubCategoryServiceInterface
	logService         service.LogServiceInterface
}

func SubCategoryController(subCategoryService service.SubCategoryServiceInterface, logService service.LogServiceInterface) *subCategoryController {
	return &subCategoryController{subCategoryService, logService}
}

func (controller *subCategoryController) GetSubCategory(context *gin.Context) {

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	sub_category, error := controller.subCategoryService.GetSubCategory()

	if error == nil {
		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":          status,
			"listSubCategory": sub_category,
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
	parse_sub_category, _ := json.Marshal(sub_category)
	var result = fmt.Sprintf("{\"status\": %s, \"listRole\": %s}", string(parse_status), string(parse_sub_category))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
