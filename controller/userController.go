package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"time"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserServiceInterface
	logService  service.LogServiceInterface
}

func UserController(userService service.UserServiceInterface, logService service.LogServiceInterface) *userController {
	return &userController{userService, logService}
}

func (controller *userController) GetUser(context *gin.Context) {

	search := context.Param("search")
	size, error := strconv.Atoi(context.Param("size"))
	page_no, error := strconv.Atoi(context.Param("page_no"))

	if search == "*" {
		search = ""
	}
	request := model.GetUserRequest{
		Search: search,
		Size:   size,
		PageNo: page_no,
	}

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	user, total_pages, error := controller.userService.GetUser(request)

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":     status,
			"listUser":   user,
			"page":       page_no,
			"totalPages": total_pages,
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
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"listUser\": %s, \"page\": %d, \"totalPages\": %d}", string(parse_status), string(parse_user), page_no, int(total_pages))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *userController) GetUserDetail(context *gin.Context) {

	username := context.Param("username")

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	user, error := controller.userService.GetUserDetail(username)

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":   status,
			"listUser": user,
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
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"listUser\": %s", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
