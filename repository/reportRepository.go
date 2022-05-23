package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type ReportRepositoryInterface interface {
	GetReport(request model.GetReportRequest) ([]entity.Ticket, error)
}

func (repo *repository) GetReport(request model.GetReportRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	var query string

	if len(request.Category) == 0 {
		query = "SELECT DISTINCT ON (ticket_code) * FROM (SELECT ticket.*, category.name AS category, ticket_isi.isi FROM ticket LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) LEFT OUTER JOIN ticket_isi ON (ticket.ticket_code = ticket_isi.ticket_code) WHERE prioritas IN @Priority AND status IN @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND ticket.tgl_dibuat >= @StartDate AND ticket.tgl_dibuat <= @EndDate ORDER BY ticket_isi.tgl_dibuat ASC) AS ticket"
	} else {
		query = "SELECT DISTINCT ON (ticket_code) * FROM (SELECT ticket.*, category.name AS category, ticket_isi.isi FROM ticket LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) LEFT OUTER JOIN ticket_isi ON (ticket.ticket_code = ticket_isi.ticket_code) WHERE prioritas IN @Priority AND status IN @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND category IN @Category AND ticket.tgl_dibuat >= @StartDate AND ticket.tgl_dibuat <= @EndDate ORDER BY ticket_isi.tgl_dibuat ASC) AS ticket"
	}

	error := repo.db.Raw(query, model.GetReportRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		Priority:        request.Priority,
		Status:          request.Status,
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error
	fmt.Println(len(ticket))
	return ticket, error
}
