package repository

import "svc-myg-ticketing/entity"

type TicketIsiRepositoryInterface interface {
	GetTicketIsi(ticket_code string) ([]entity.TicketIsi, error)
}

func (repo *repository) GetTicketIsi(ticket_code string) ([]entity.TicketIsi, error) {
	var ticket_isi []entity.TicketIsi

	error := repo.db.Raw("SELECT * FROM ticket_isi WHERE kode_ticket = ?", ticket_code).Find(&ticket_isi).Error

	return ticket_isi, error
}
