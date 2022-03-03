package service

import (
	"math"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type TicketServiceInterface interface {
	GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error)
}

type ticketService struct {
	repository repository.TicketRepositoryInterface
}

func TicketService(repository repository.TicketRepositoryInterface) *ticketService {
	return &ticketService{repository}
}

func (ticketService *ticketService) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error) {

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := ticketService.repository.CountTicket(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))

	ticket, error := ticketService.repository.GetTicket(request)

	return ticket, int(total_pages), error
}
