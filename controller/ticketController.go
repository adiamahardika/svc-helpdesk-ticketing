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

type ticketController struct {
	ticketService service.TicketServiceInterface
	logService    service.LogServiceInterface
}

func TicketController(ticketService service.TicketServiceInterface, logService service.LogServiceInterface) *ticketController {
	return &ticketController{ticketService, logService}
}

func (controller *ticketController) GetTicket(context *gin.Context) {

	var request *model.GetTicketRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var ticket []entity.Ticket
	var total_pages int

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

		ticket, total_pages, error = controller.ticketService.GetTicket(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":     status,
				"content":    ticket,
				"page":       request.PageNo,
				"totalPages": total_pages,
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
	parse_ticket, _ := json.Marshal(ticket)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s, \"page\": %d, \"totalPages\": %d}", string(parse_status), string(parse_ticket), request.PageNo, total_pages)
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *ticketController) GetDetailTicket(context *gin.Context) {

	ticket_code := context.Param("ticket-code")

	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	ticket, error := controller.ticketService.GetDetailTicket(&ticket_code)

	if error == nil {

		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatusCode: http.StatusOK,
			ResponseCode:   general.SuccessStatusCode,
			Description:    description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":           status,
			"listDetailTicket": ticket.ListDetailTicket,
			"listReplyTicket":  ticket.ListReplyTicket,
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
	parse_ticket, _ := json.Marshal(ticket.ListDetailTicket)
	parse_reply, _ := json.Marshal(ticket.ListReplyTicket)
	var result = fmt.Sprintf("{\"status\": %s, \"listDetailTicket\": %s, \"listReplyTicket\": %s}", string(parse_status), string(parse_ticket), string(parse_reply))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *ticketController) CreateTicket(context *gin.Context) {

	var request *model.CreateTicketRequest

	error := context.ShouldBind(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var ticket entity.Ticket
	var ticket_isi entity.TicketIsi

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

		ticket, ticket_isi, error = controller.ticketService.CreateTicket(request, context)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":     status,
				"ticket":     ticket,
				"ticket_isi": ticket_isi,
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
	parse_ticket, _ := json.Marshal(ticket)
	parse_ticket_isi, _ := json.Marshal(ticket_isi)
	var result = fmt.Sprintf("{\"status\": %s, \"ticket\": %s, \"ticket_isi\": %s}", string(parse_status), string(parse_ticket), string(parse_ticket_isi))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *ticketController) UpdateTicket(context *gin.Context) {

	var request *model.UpdateTicketRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var ticket []entity.Ticket

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

		ticket, error = controller.ticketService.UpdateTicket(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"ticket": ticket,
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
	parse_ticket, _ := json.Marshal(ticket)
	var result = fmt.Sprintf("{\"status\": %s, \"ticket\": %s}", string(parse_status), string(parse_ticket))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *ticketController) ReplyTicket(context *gin.Context) {

	var request *model.ReplyTicket

	error := context.ShouldBind(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var ticket []entity.Ticket

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

		ticket, error = controller.ticketService.ReplyTicket(request, context)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"ticket": ticket,
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
	parse_ticket, _ := json.Marshal(ticket)
	var result = fmt.Sprintf("{\"status\": %s, \"ticket\": %s}", string(parse_status), string(parse_ticket))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *ticketController) UpdateTicketStatus(context *gin.Context) {

	var request *model.UpdateTicketStatusRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var ticket []entity.Ticket

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

		ticket, error = controller.ticketService.UpdateTicketStatus(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"ticket": ticket,
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
	parse_ticket, _ := json.Marshal(ticket)
	var result = fmt.Sprintf("{\"status\": %s, \"ticket\": %s}", string(parse_status), string(parse_ticket))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
