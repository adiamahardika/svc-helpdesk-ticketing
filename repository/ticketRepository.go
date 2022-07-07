package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type TicketRepositoryInterface interface {
	GetTicket(request *model.GetTicketRequest) ([]*entity.Ticket, error)
	CountTicket(request *model.GetTicketRequest) (*int, error)
	GetDetailTicket(ticket_code *string) (*entity.Ticket, error)
	CreateTicket(request *entity.Ticket) (*entity.Ticket, error)
	CheckTicketCode(request *string) ([]*entity.Ticket, error)
	UpdateTicket(request *model.UpdateTicketRequest) ([]*entity.Ticket, error)
	UpdateTicketStatus(request *model.UpdateTicketStatusRequest) ([]*entity.Ticket, error)
}

func (repo *repository) GetTicket(request *model.GetTicketRequest) ([]*entity.Ticket, error) {
	var ticket []*entity.Ticket
	var query string

	if len(request.Category) == 0 {
		query = "SELECT * FROM (SELECT ticket.*, ticketing_category.name AS category_name, ms_area.area_name, ms_grapari.name AS grapari_name FROM ticket LEFT OUTER JOIN ticketing_category ON (ticket.category = CAST(ticketing_category.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) WHERE prioritas LIKE @Priority AND ticket.status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
	} else {
		query = "SELECT * FROM (SELECT ticket.*, ticketing_category.name AS category_name, ms_area.area_name, ms_grapari.name AS grapari_name FROM ticket LEFT OUTER JOIN ticketing_category ON (ticket.category = CAST(ticketing_category.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) WHERE prioritas LIKE @Priority AND ticket.status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND category IN @Category AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex"
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
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) CountTicket(request *model.GetTicketRequest) (*int, error) {
	var total_data *int
	var query string

	if len(request.Category) == 0 {
		query = "SELECT COUNT(*) as total_data FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search"
	} else {
		query = "SELECT COUNT(*) as total_data FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat AND category IN @Category AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search"
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
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) GetDetailTicket(ticket_code *string) (*entity.Ticket, error) {
	var ticket *entity.Ticket

	error := repo.db.Raw("SELECT ticket.*, ms_area.area_name, ms_grapari.name AS grapari_name, ticketing_category.name AS category_name FROM ticket LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN ticketing_category ON (ticket.category = CAST(ticketing_category.id AS varchar(10))) WHERE ticket_code = ?", ticket_code).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) CreateTicket(request *entity.Ticket) (*entity.Ticket, error) {
	var ticket *entity.Ticket

	error := repo.db.Table("ticket").Create(&request).Error

	return ticket, error
}

func (repo *repository) CheckTicketCode(request *string) ([]*entity.Ticket, error) {
	var ticket []*entity.Ticket

	error := repo.db.Raw("SELECT ticket.* FROM ticket WHERE ticket_code = @TicketCode", model.CreateTicketRequest{
		TicketCode: *request,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) UpdateTicket(request *model.UpdateTicketRequest) ([]*entity.Ticket, error) {
	var ticket []*entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET assigned_to =  @AssignedTo, email = @Email, category = @Category, sub_category = @SubCategory,  prioritas = @Prioritas, status = @Status, username_pembalas = @UsernamePembalas, tgl_diperbarui = @TglDiperbarui WHERE ticket_code = @TicketCode RETURNING ticket.*", request).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) UpdateTicketStatus(request *model.UpdateTicketStatusRequest) ([]*entity.Ticket, error) {
	var ticket []*entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET status = @Status, tgl_diperbarui = @TglDiperbarui WHERE ticket_code = @TicketCode RETURNING ticket.*", request).Find(&ticket).Error

	return ticket, error
}
