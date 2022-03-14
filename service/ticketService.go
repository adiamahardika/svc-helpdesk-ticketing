package service

import (
	"fmt"
	"math"
	"os"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketServiceInterface interface {
	GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error)
	GetDetailTicket(ticket_code string) (model.GetDetailTicketResponse, error)
	CreateTicket(request model.CreateTicketRequest, context *gin.Context) (entity.Ticket, entity.TicketIsi, error)
}

type ticketService struct {
	ticketRepository    repository.TicketRepositoryInterface
	ticketIsiRepository repository.TicketIsiRepositoryInterface
}

func TicketService(ticketRepository repository.TicketRepositoryInterface, ticketIsiRepository repository.TicketIsiRepositoryInterface) *ticketService {
	return &ticketService{ticketRepository, ticketIsiRepository}
}

func (ticketService *ticketService) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, int, error) {

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := ticketService.ticketRepository.CountTicket(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))

	ticket, error := ticketService.ticketRepository.GetTicket(request)

	return ticket, int(total_pages), error
}

func (ticketService *ticketService) GetDetailTicket(ticket_code string) (model.GetDetailTicketResponse, error) {

	var reponse model.GetDetailTicketResponse

	detail_ticket, error := ticketService.ticketRepository.GetDetailTicket(ticket_code)
	reply_ticket, error := ticketService.ticketIsiRepository.GetTicketIsi(ticket_code)

	url := os.Getenv("FILE_URL")

	for index := range reply_ticket {
		date := reply_ticket[index].TglDibuat.Format("2006-01-02")
		ticket_code := reply_ticket[index].TicketCode

		file_name1 := reply_ticket[index].Attachment1
		if file_name1 != "" {
			reply_ticket[index].Attachment1 = url + "ticket/" + ticket_code + "/" + date + "/" + file_name1
		}

		file_name2 := reply_ticket[index].Attachment2
		if file_name2 != "" {
			reply_ticket[index].Attachment2 = url + "ticket/" + ticket_code + "/" + date + "/" + file_name2
		}
	}

	reponse.ListDetailTicket = detail_ticket
	reponse.ListReplyTicket = reply_ticket

	return reponse, error
}

func (ticketService *ticketService) CreateTicket(request model.CreateTicketRequest, context *gin.Context) (entity.Ticket, entity.TicketIsi, error) {
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

	ticket_request := entity.Ticket{
		Judul:            request.Judul,
		UsernamePembuat:  request.UserPembuat,
		UsernamePembalas: request.UserPembuat,
		Prioritas:        request.Prioritas,
		TotalWaktu:       request.TotalWaktu,
		Status:           request.Status,
		TicketCode:       request.TicketCode,
		Category:         request.Category,
		Lokasi:           request.Lokasi,
		TerminalId:       request.TerminalId,
		Email:            request.Email,
		AssignedTo:       request.AssignedTo,
		TglDibuat:        date_now,
		TglDiperbarui:    date_now,
	}

	ticket_isi_request := entity.TicketIsi{
		UsernamePengirim: request.UserPembuat,
		Isi:              request.Isi,
		TicketCode:       request.TicketCode,
		Attachment1:      attachment1,
		Attachment2:      attachment2,
		TglDibuat:        date_now,
	}

	ticket, error = ticketService.ticketRepository.CheckTicketCode(request.TicketCode)

	if len(ticket) > 0 {
		error = fmt.Errorf("Ticket code already exist!")
	} else if error == nil {

		_, error = ticketService.ticketRepository.CreateTicket(ticket_request)

		if error == nil {
			_, error = ticketService.ticketIsiRepository.CreateTicketIsi(ticket_isi_request)
		}
	}

	return ticket_request, ticket_isi_request, error

}
