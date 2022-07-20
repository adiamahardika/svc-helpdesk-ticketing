package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
)

type TicketStatusServiceInterface interface {
	GetTicketStatus() ([]entity.TicketStatus, error)
}

type ticketStatusService struct {
	ticketStatusRepository repository.TicketStatusRepositoryInterface
}

func TicketStatusService(ticketStatusRepository repository.TicketStatusRepositoryInterface) *ticketStatusService {
	return &ticketStatusService{ticketStatusRepository}
}

func (ticketStatusService *ticketStatusService) GetTicketStatus() ([]entity.TicketStatus, error) {

	ticket_status, error := ticketStatusService.ticketStatusRepository.GetTicketStatus()

	return ticket_status, error
}
