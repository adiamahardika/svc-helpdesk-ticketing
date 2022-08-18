package repository

import (
	"fmt"
	"svc-myg-ticketing/model"
)

type ReportRepositoryInterface interface {
	GetReport(request *model.GetReportRequest) ([]model.ReportResponse, error)
}

func (repo *repository) GetReport(request *model.GetReportRequest) ([]model.ReportResponse, error) {
	var ticket []model.ReportResponse
	var query string
	var category string
	var created_by string
	var area_code string
	var regional string
	var grapari_id string

	if len(request.Category) > 0 {
		category = "AND ticketing_category IN @Category"
	}
	if len(request.UsernamePembuat) > 0 {
		created_by = "AND username_pembuat IN @UsernamePembuat"
	}
	if len(request.AreaCode) > 0 {
		area_code = "AND ticket.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND ticket.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND ticket.grapari_id IN @GrapariId"
	}

	query = fmt.Sprintf("SELECT DISTINCT ON (ticket_code) ticket.*, TO_CHAR(ticket.tgl_diperbarui, 'DD-MM-YYYY HH24:MI:SS') AS tgl_diperbarui, TO_CHAR(ticket.tgl_dibuat, 'DD-MM-YYYY HH24:MI:SS' ) AS tgl_dibuat, TO_CHAR(ticket.start_time, 'DD-MM-YYYY HH24:MI:SS' ) AS start_time, TO_CHAR(ticket.close_time, 'DD-MM-YYYY HH24:MI:SS' ) AS close_time, TO_CHAR(ticket.assigning_time, 'DD-MM-YYYY HH24:MI:SS' ) AS assigning_time, ticketing_category.name AS category, ticket_isi.isi, ms_area.area_name, ms_grapari.name AS grapari_name, users1.name AS user_pembuat, users2.name AS assignee FROM ticket LEFT OUTER JOIN ticketing_category ON (ticket.category = CAST(ticketing_category.id AS varchar(10))) LEFT OUTER JOIN ticket_isi ON (ticket.ticket_code = ticket_isi.ticket_code) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users1 ON (ticket.username_pembuat = users1.username) LEFT OUTER JOIN users users2 ON (ticket.assigned_to = users2.username) WHERE prioritas IN @Priority AND ticket.status IN @Status AND assigned_to LIKE @AssignedTo %s %s %s %s %s AND ticket.tgl_dibuat >= @StartDate AND ticket.tgl_dibuat <= @EndDate ORDER BY ticket_code ASC", category, created_by, area_code, regional, grapari_id)

	error := repo.db.Raw(query, model.GetReportRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		AreaCode:        request.AreaCode,
		Regional:        request.Regional,
		GrapariId:       request.GrapariId,
		Priority:        request.Priority,
		Status:          request.Status,
		UsernamePembuat: request.UsernamePembuat,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}
