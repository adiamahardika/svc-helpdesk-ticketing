package repository

import "svc-myg-ticketing/entity"

type TicketIsiRepositoryInterface interface {
	GetTicketIsi(ticket_code *string) ([]entity.TicketIsi, error)
	CreateTicketIsi(request *entity.TicketIsi) (entity.TicketIsi, error)
}

func (repo *repository) GetTicketIsi(ticket_code *string) ([]entity.TicketIsi, error) {
	var ticket_isi []entity.TicketIsi

	error := repo.db.Raw("SELECT ticket_isi.*, users.name AS username_pengirim FROM ticket_isi LEFT OUTER JOIN users ON (ticket_isi.username_pengirim = users.username) WHERE ticket_code = ? ORDER BY tgl_dibuat ASC", ticket_code).Find(&ticket_isi).Error

	return ticket_isi, error
}

func (repo *repository) CreateTicketIsi(request *entity.TicketIsi) (entity.TicketIsi, error) {
	var ticket_isi entity.TicketIsi

	error := repo.db.Table("ticket_isi").Create(&request).Error

	return ticket_isi, error
}
