package repository

import (
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
		query = "SELECT ticket.*, category.name AS category FROM ticket LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) WHERE prioritas IN @Priority AND status IN @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_dibuat ASC"
	} else {
		query = "SELECT ticket.*, category.name AS category FROM ticket LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) WHERE prioritas IN @Priority AND status IN @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND category IN @Category AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_dibuat ASC"
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

	return ticket, error
}
