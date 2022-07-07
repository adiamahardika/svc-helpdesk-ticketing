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

type ticketStatusController struct {
	ticketStatusService service.TicketStatusServiceInterface
	logService          service.LogServiceInterface
}

func TicketStatusController(ticketStatusService service.TicketStatusServiceInterface, logService service.LogServiceInterface) *ticketStatusController {
	return &ticketStatusController{ticketStatusService, logService}
}

func (controller *ticketStatusController) GetTicketStatus(context *gin.Context) {

	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	ticket_status, error := controller.ticketStatusService.GetTicketStatus()

	if error == nil {

		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": ticket_status,
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
	parse_result, _ := json.Marshal(ticket_status)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_result))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
