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

type categoryController struct {
	categoryService service.CategoryServiceInterface
	logService      service.LogServiceInterface
}

func CategoryController(categoryService service.CategoryServiceInterface, logService service.LogServiceInterface) *categoryController {
	return &categoryController{categoryService, logService}
}

func (controller *categoryController) GetCategory(context *gin.Context) {

	size := context.Param("size")
	page_no := context.Param("page_no")
	sort_by := context.Param("sort_by")
	order_by := context.Param("order_by")

	request := model.GetCategoryRequest{
		Size:     size,
		PageNo:   page_no,
		SortBy:   sort_by,
		OrderBy:  order_by,
		IsActive: "true",
	}

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	category, error := controller.categoryService.GetCategory(request)

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": category,
		})
	} else {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	}

	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_category, _ := json.Marshal(category)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
