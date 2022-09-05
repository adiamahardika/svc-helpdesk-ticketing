package repository

import (
	"fmt"
	"svc-myg-ticketing/model"
)

type ReportRepositoryInterface interface {
	GetReport(request *model.GetReportRequest) ([]model.ReportResponse, error)
	GetCountReportByStatus(request *model.GetCountReportByStatusRequest) ([]model.GetCountReportByStatusResponse, error)
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

func (repo *repository) GetCountReportByStatus(request *model.GetCountReportByStatusRequest) ([]model.GetCountReportByStatusResponse, error) {

	var result []model.GetCountReportByStatusResponse
	var area_code string
	var regional string
	var grapari_id string

	if len(request.AreaCode) > 0 {
		area_code = "AND ticket.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND ticket.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND ticket.grapari_id IN @GrapariId"
	}

	query1 := fmt.Sprintf("(SELECT * FROM (SELECT DATE(tgl_dibuat) AS date, COUNT(*) As new, 0 AS process, 0 As finish FROM myg_ticketing.ticket WHERE tgl_dibuat >= @StartDate and tgl_dibuat <= @EndDate  %s %s %s GROUP BY DATE(tgl_dibuat)) as tbl)", area_code, regional, grapari_id)
	query2 := fmt.Sprintf("(SELECT * FROM (SELECT DATE(start_time) AS date, 0 AS new, COUNT(*) As process, 0 As finish FROM myg_ticketing.ticket WHERE start_time >= @StartDate and start_time <= @EndDate  %s %s %s GROUP BY DATE(start_time)) as tbl)", area_code, regional, grapari_id)
	query3 := fmt.Sprintf("(SELECT * FROM (SELECT DATE(close_time) AS date, 0 AS new, 0 As process, COUNT(*) As finish FROM myg_ticketing.ticket WHERE close_time >= @StartDate and close_time <= @EndDate  %s %s %s GROUP BY DATE(close_time)) as tbl)", area_code, regional, grapari_id)

	final_query := fmt.Sprintf("SELECT date, SUM(new) AS new, SUM(process) AS process, SUM(finish) AS finish FROM (%s UNION ALL %s UNION ALL %s) AS tbl2 GROUP BY date ORDER BY date ASC", query1, query2, query3)

	error := repo.db.Raw(final_query, model.GetCountReportByStatusRequest{
		AreaCode:  request.AreaCode,
		Regional:  request.Regional,
		GrapariId: request.GrapariId,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}).Find(&result).Error

	return result, error
}
