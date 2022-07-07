package repository

import "svc-myg-ticketing/entity"

type TicketStatusRepositoryInterface interface {
	GetTicketStatus() ([]*entity.TicketStatus, error)
}

func (repo *repository) GetTicketStatus() ([]*entity.TicketStatus, error) {
	var ticket_status []*entity.TicketStatus

	error := repo.db.Raw("SELECT * FROM ticket_status ORDER BY index ASC").Find(&ticket_status).Error

	return ticket_status, error
}
