package service

import (
	"math"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type TicketServiceInterface interface {
	GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error)
	GetDetailTicket(ticket_code string) (model.GetDetailTicketResponse, error)
}

type ticketService struct {
	ticketRepository    repository.TicketRepositoryInterface
	ticketIsiRepository repository.TicketIsiRepositoryInterface
}

func TicketService(ticketRepository repository.TicketRepositoryInterface, ticketIsiRepository repository.TicketIsiRepositoryInterface) *ticketService {
	return &ticketService{ticketRepository, ticketIsiRepository}
}

func (ticketService *ticketService) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error) {

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := ticketService.ticketRepository.CountTicket(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))

	ticket, error := ticketService.ticketRepository.GetTicket(request)

	return ticket, int(total_pages), error
}

func (ticketService *ticketService) GetDetailTicket(ticket_code string) (model.GetDetailTicketResponse, error) {

	var reponse model.GetDetailTicketResponse

	detail_ticket, error := ticketService.ticketRepository.GetDetailTicket(ticket_code)
	reply_ticket, error := ticketService.ticketIsiRepository.GetTicketIsi(ticket_code)

	reponse.ListDetailTicket = detail_ticket
	reponse.ListReplyTicket = reply_ticket

	return reponse, error
}
