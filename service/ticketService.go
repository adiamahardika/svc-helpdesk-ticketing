package service

import (
	"fmt"
	"math"
	"os"
	"strings"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketServiceInterface interface {
	GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, int, error)
	GetDetailTicket(ticket_code *string) (model.GetDetailTicketResponse, error)
	CreateTicket(request *model.CreateTicketRequest, context *gin.Context) (entity.Ticket, entity.TicketIsi, error)
	UpdateTicket(request *model.UpdateTicketRequest) ([]entity.Ticket, error)
	ReplyTicket(request *model.ReplyTicket, context *gin.Context) ([]entity.Ticket, error)
	UpdateTicketStatus(request *model.UpdateTicketStatusRequest) ([]entity.Ticket, error)
	StartTicket(request *model.StartTicketRequest) ([]entity.Ticket, error)
	CloseTicket(request *model.CloseTicketRequest) ([]entity.Ticket, error)
}

type ticketService struct {
	ticketRepository     repository.TicketRepositoryInterface
	ticketIsiRepository  repository.TicketIsiRepositoryInterface
	emailNotifRepository repository.EmailNotifRepositoryInterface
}

func TicketService(ticketRepository repository.TicketRepositoryInterface, ticketIsiRepository repository.TicketIsiRepositoryInterface, emailNotifRepository repository.EmailNotifRepositoryInterface) *ticketService {
	return &ticketService{ticketRepository, ticketIsiRepository, emailNotifRepository}
}

func (ticketService *ticketService) GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, int, error) {

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}
	request.EndDate = request.EndDate + " 23:59:59"
	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := ticketService.ticketRepository.CountTicket(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))

	ticket, error := ticketService.ticketRepository.GetTicket(request)
	parse_tp := int(total_pages)

	return ticket, parse_tp, error
}

func (ticketService *ticketService) GetDetailTicket(ticket_code *string) (model.GetDetailTicketResponse, error) {

	var reponse model.GetDetailTicketResponse

	detail_ticket, error := ticketService.ticketRepository.GetDetailTicket(ticket_code)
	reply_ticket, error := ticketService.ticketIsiRepository.GetTicketIsi(ticket_code)

	url := os.Getenv("FILE_URL")

	for index := range reply_ticket {
		date := reply_ticket[index].TglDibuat.Format("2006-01-02")
		ticket_code := reply_ticket[index].TicketCode

		file_name1 := reply_ticket[index].Attachment1
		if file_name1 != "-" {
			reply_ticket[index].Attachment1 = url + "ticket/" + ticket_code + "/" + date + "/" + file_name1
		}

		file_name2 := reply_ticket[index].Attachment2
		if file_name2 != "-" {
			reply_ticket[index].Attachment2 = url + "ticket/" + ticket_code + "/" + date + "/" + file_name2
		}
	}

	reponse.ListDetailTicket = detail_ticket
	reponse.ListReplyTicket = reply_ticket

	return reponse, error
}

func (ticketService *ticketService) CreateTicket(request *model.CreateTicketRequest, context *gin.Context) (entity.Ticket, entity.TicketIsi, error) {
	var wg sync.WaitGroup
	var ticket []entity.Ticket
	var assigning_time time.Time

	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "/ticket/" + request.TicketCode + "/" + date_now.Format("2006-01-02")
	error := fmt.Errorf("error")
	attachment1 := "-"
	attachment2 := "-"
	assigning_by := request.UserPembuat

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error = os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if request.Attachment1 != nil {
		attachment1 = general.RandomString(4) + "_" + request.Attachment1.Filename
		error = context.SaveUploadedFile(request.Attachment1, path+"/"+attachment1)
	} else {
		error = nil
	}

	if request.Attachment2 != nil {
		attachment2 = general.RandomString(4) + "_" + request.Attachment2.Filename
		error = context.SaveUploadedFile(request.Attachment2, path+"/"+attachment2)
	} else {
		error = nil
	}
	total_waktu := "0y 0m 0d 0h 0mm 0s"
	if request.AssignedTo == "Unassigned" {
		assigning_time = time.Time{}
		assigning_by = ""
	} else {
		assigning_time = date_now
	}

	ticket_request := entity.Ticket{
		Judul:             request.Judul,
		UsernamePembuat:   request.UserPembuat,
		UpdatedBy:         request.UserPembuat,
		Prioritas:         request.Prioritas,
		TotalWaktu:        total_waktu,
		Status:            request.Status,
		TicketCode:        request.TicketCode,
		Category:          request.Category,
		SubCategory:       request.SubCategory,
		Lokasi:            request.Lokasi,
		TerminalId:        request.TerminalId,
		AreaCode:          request.AreaCode,
		Regional:          request.Regional,
		GrapariId:         request.GrapariId,
		Email:             request.Email,
		AssignedTo:        request.AssignedTo,
		EmailNotification: request.EmailNotification,
		TglDibuat:         date_now,
		TglDiperbarui:     date_now,
		AssigningTime:     assigning_time,
		AssigningBy:       assigning_by,
	}

	ticket_isi_request := entity.TicketIsi{
		UsernamePengirim: request.UserPembuat,
		Isi:              request.Isi,
		TicketCode:       request.TicketCode,
		Attachment1:      attachment1,
		Attachment2:      attachment2,
		TglDibuat:        date_now,
	}

	ticket, error = ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) > 0 {
		error = fmt.Errorf("Ticket code already exist!")
	} else if error == nil {

		_, error = ticketService.ticketRepository.CreateTicket(&ticket_request)

		if error == nil {
			_, error = ticketService.ticketIsiRepository.CreateTicketIsi(&ticket_isi_request)
		}

		if request.EmailNotification == "true" {
			// wg.Add(1)
			var detail_ticket entity.Ticket
			detail_ticket, error = ticketService.ticketRepository.GetDetailTicket(&request.TicketCode)
			assignee := detail_ticket.Assignee

			if len(assignee) == 0 {
				assignee = "Unassigned"
			}

			email_notif, _ := ticketService.emailNotifRepository.GetAllEmailNotif()
			sender := NewSMTP()
			message := NewMessage(&model.SmtpRequest{
				Judul:           request.Judul,
				Prioritas:       request.Prioritas,
				UsernamePembuat: request.UserPembuat,
				Status:          request.Status,
				TicketCode:      request.TicketCode,
				Lokasi:          request.Lokasi,
				TerminalId:      request.TerminalId,
				Email:           request.Email,
				Isi:             request.Isi,
				AreaName:        detail_ticket.AreaName,
				Regional:        detail_ticket.Regional,
				GrapariName:     detail_ticket.GrapariName,
				CategoryName:    detail_ticket.CategoryName,
				SubCategory:     detail_ticket.SubCategory,
				UserPembuat:     detail_ticket.UserPembuat,
				Assignee:        assignee,
				Type:            "New",
			})
			message.CC = []string{request.Email}
			message.To = email_notif
			message.AttachFile(path+attachment1, path+attachment2)
			sender.Send(&wg, message)
		}

	}

	return ticket_request, ticket_isi_request, error

}

func (ticketService *ticketService) UpdateTicket(request *model.UpdateTicketRequest) ([]entity.Ticket, error) {
	date_now := time.Now()

	ticket, error := ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) == 0 {
		error = fmt.Errorf("Ticket code does'nt exist!")
	} else {

		request.TglDiperbarui = date_now
		request.AssigningTime = ticket[0].AssigningTime
		request.AssigningBy = ticket[0].AssigningBy

		if request.AssignedTo != ticket[0].AssignedTo {
			request.AssigningTime = date_now
			request.AssigningBy = request.UpdatedBy
		}

		if ticket[0].Status != "Process" && request.Status == "Process" {
			start_req := model.StartTicketRequest{
				TicketCode: request.TicketCode,
				StartTime:  date_now,
				StartBy:    request.UpdatedBy,
			}
			_, error = ticketService.ticketRepository.StartTicket(&start_req)

			close_req := model.CloseTicketRequest{
				TicketCode: request.TicketCode,
				CloseTime:  time.Time{},
				CloseBy:    request.UpdatedBy,
			}
			_, error = ticketService.ticketRepository.CloseTicket(&close_req)
		} else if ticket[0].Status != "Finish" && request.Status == "Finish" {
			close_req := model.CloseTicketRequest{
				TicketCode: request.TicketCode,
				CloseTime:  date_now,
				CloseBy:    request.UpdatedBy,
			}

			_, error = ticketService.ticketRepository.CloseTicket(&close_req)
		} else if ticket[0].Status != request.Status {
			close_req := model.CloseTicketRequest{
				TicketCode: request.TicketCode,
				CloseTime:  time.Time{},
				CloseBy:    "",
			}

			_, error = ticketService.ticketRepository.CloseTicket(&close_req)
		}

		ticket, error = ticketService.ticketRepository.UpdateTicket(request)
	}

	return ticket, error
}

func (ticketService *ticketService) ReplyTicket(request *model.ReplyTicket, context *gin.Context) ([]entity.Ticket, error) {
	var ticket []entity.Ticket

	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "/ticket/" + request.TicketCode + "/" + date_now.Format("2006-01-02")
	error := fmt.Errorf("error")
	attachment1 := "-"
	attachment2 := "-"

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error = os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if request.Attachment1 != nil {
		attachment1 = general.RandomString(4) + "_" + request.Attachment1.Filename
		error = context.SaveUploadedFile(request.Attachment1, path+"/"+attachment1)
	} else {
		error = nil
	}

	if request.Attachment2 != nil {
		attachment2 = general.RandomString(4) + "_" + request.Attachment2.Filename
		error = context.SaveUploadedFile(request.Attachment2, path+"/"+attachment2)
	} else {
		error = nil
	}

	ticket, error = ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) == 0 {
		error = fmt.Errorf("Ticket code does'nt exist!")
	} else {

		if strings.EqualFold(request.ReplyType, "start") {
			start_req := model.StartTicketRequest{
				TicketCode: request.TicketCode,
				StartTime:  date_now,
				StartBy:    request.UsernamePengirim,
			}

			_, error = ticketService.ticketRepository.StartTicket(&start_req)

		} else if strings.EqualFold(request.ReplyType, "close") {
			close_req := model.CloseTicketRequest{
				TicketCode: request.TicketCode,
				CloseTime:  date_now,
				CloseBy:    request.UsernamePengirim,
			}

			_, error = ticketService.ticketRepository.CloseTicket(&close_req)
		}

		update_ticket := model.UpdateTicketRequest{
			AssignedTo:    ticket[0].AssignedTo,
			Category:      ticket[0].Category,
			Prioritas:     ticket[0].Prioritas,
			SubCategory:   ticket[0].SubCategory,
			AssigningTime: ticket[0].AssigningTime,
			AssigningBy:   ticket[0].AssigningBy,
			VisitStatus:   request.VisitStatus,
			Status:        request.Status,
			TicketCode:    request.TicketCode,
			UpdatedBy:     request.UpdatedBy,
			TglDiperbarui: date_now,
		}
		if error == nil {
			_, error = ticketService.ticketRepository.UpdateTicket(&update_ticket)
		}

		reply_request := entity.TicketIsi{
			TicketCode:       request.TicketCode,
			UsernamePengirim: request.UsernamePengirim,
			Isi:              request.Isi,
			Attachment1:      attachment1,
			Attachment2:      attachment2,
			TglDibuat:        date_now,
		}
		if error == nil {
			_, error = ticketService.ticketIsiRepository.CreateTicketIsi(&reply_request)
		}

	}

	return ticket, error
}

func (ticketService *ticketService) UpdateTicketStatus(request *model.UpdateTicketStatusRequest) ([]entity.Ticket, error) {
	date_now := time.Now()

	ticket, error := ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) == 0 {
		error = fmt.Errorf("Ticket code does'nt exist!")
	} else {

		request.TglDiperbarui = date_now

		ticket, error = ticketService.ticketRepository.UpdateTicketStatus(request)
	}

	return ticket, error
}

func (ticketService *ticketService) StartTicket(request *model.StartTicketRequest) ([]entity.Ticket, error) {

	ticket, error := ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) == 0 {
		error = fmt.Errorf("Ticket code doesnt exist!")
	} else {
		date_now := time.Now()
		update_ticket := model.UpdateTicketRequest{
			AssignedTo:    ticket[0].AssignedTo,
			Category:      ticket[0].Category,
			Prioritas:     ticket[0].Prioritas,
			SubCategory:   ticket[0].SubCategory,
			AssigningTime: ticket[0].AssigningTime,
			AssigningBy:   ticket[0].AssigningBy,
			VisitStatus:   ticket[0].VisitStatus,
			Status:        "Process",
			TicketCode:    request.TicketCode,
			UpdatedBy:     request.StartBy,
			TglDiperbarui: date_now,
		}
		if error == nil {
			_, error = ticketService.ticketRepository.UpdateTicket(&update_ticket)
		}

		request.StartTime = date_now
		if error == nil {
			ticket, error = ticketService.ticketRepository.StartTicket(request)
		}
	}
	return ticket, error
}

func (ticketService *ticketService) CloseTicket(request *model.CloseTicketRequest) ([]entity.Ticket, error) {

	ticket, error := ticketService.ticketRepository.CheckTicketCode(&request.TicketCode)

	if len(ticket) == 0 {
		error = fmt.Errorf("Ticket code does'nt exist!")
	} else {
		date_now := time.Now()
		update_ticket := model.UpdateTicketRequest{
			AssignedTo:    ticket[0].AssignedTo,
			Category:      ticket[0].Category,
			Prioritas:     ticket[0].Prioritas,
			SubCategory:   ticket[0].SubCategory,
			AssigningTime: ticket[0].AssigningTime,
			AssigningBy:   ticket[0].AssigningBy,
			VisitStatus:   request.VisitStatus,
			Status:        "Finish",
			TicketCode:    request.TicketCode,
			UpdatedBy:     request.CloseBy,
			TglDiperbarui: date_now,
		}
		if error == nil {
			_, error = ticketService.ticketRepository.UpdateTicket(&update_ticket)
		}

		request.CloseTime = date_now
		if error == nil {
			ticket, error = ticketService.ticketRepository.CloseTicket(request)
		}
	}
	return ticket, error
}
