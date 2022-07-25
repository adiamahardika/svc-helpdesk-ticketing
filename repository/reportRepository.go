package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type ReportRepositoryInterface interface {
	GetReport(request *model.GetReportRequest) ([]entity.Ticket, error)
}

func (repo *repository) GetReport(request *model.GetReportRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	var query string
	var category string
	var created_by string

	if len(request.Category) == 0 {
		category = ""
	} else {
		category = "AND ticketing_category IN @Category"
	}
	if len(request.UsernamePembuat) == 0 {
		created_by = ""
	} else {
		created_by = "AND username_pembuat IN @UsernamePembuat"
	}
	query = fmt.Sprintf("SELECT DISTINCT ON (ticket_code) ticket.*, ticketing_category.name AS category, ticket_isi.isi, ms_area.area_name, ms_grapari.name AS grapari_name, users1.name AS user_pembuat, users2.name AS assignee FROM ticket LEFT OUTER JOIN ticketing_category ON (ticket.category = CAST(ticketing_category.id AS varchar(10))) LEFT OUTER JOIN ticket_isi ON (ticket.ticket_code = ticket_isi.ticket_code) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users1 ON (ticket.username_pembuat = users1.username) LEFT OUTER JOIN users users2 ON (ticket.assigned_to = users2.username) WHERE prioritas IN @Priority AND ticket.status IN @Status AND assigned_to LIKE @AssignedTo %s %s AND ticket.tgl_dibuat >= @StartDate AND ticket.tgl_dibuat <= @EndDate ORDER BY ticket_code, ticket_isi.tgl_dibuat ASC", category, created_by)

	error := repo.db.Raw(query, model.GetReportRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		Priority:        request.Priority,
		Status:          request.Status,
		UsernamePembuat: request.UsernamePembuat,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}
