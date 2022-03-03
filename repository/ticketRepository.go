package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type TicketRepositoryInterface interface {
	GetTicket(request model.GetTicketRequest) ([]entity.Ticket, error)
	CountTicket(request model.GetTicketRequest) (int, error)
}

func (repo *repository) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	var query string

	if len(request.Category) == 0 {
		query = "SELECT * FROM (SELECT ticket.*, category.name AS kategori FROM ticket LEFT OUTER JOIN category ON (ticket.kategori = CAST(category.id AS varchar(10))) WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
	} else {
		query = "SELECT * FROM (SELECT ticket.*, category.name AS kategori FROM ticket LEFT OUTER JOIN category ON (ticket.kategori = CAST(category.id AS varchar(10))) WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND kategori IN @Category ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
	}

	error := repo.db.Raw(query, model.GetTicketRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		Priority:        "%" + request.Priority + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartIndex:      request.StartIndex,
		PageSize:        request.PageSize,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) CountTicket(request model.GetTicketRequest) (int, error) {
	var total_data int
	var query string

	if len(request.Category) == 0 {
		query = "SELECT COUNT(*) as total_data FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
	} else {
		query = "SELECT COUNT(*) as total_data FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND kategori IN @Category ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
	}

	error := repo.db.Raw(query, model.GetTicketRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		Priority:        "%" + request.Priority + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartIndex:      request.StartIndex,
		PageSize:        request.PageSize,
	}).Find(&total_data).Error

	return total_data, error
}
