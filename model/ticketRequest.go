package model

import (
	"mime/multipart"
	"time"
)

type GetTicketRequest struct {
	AssignedTo      string   `json:"assignedTo"`
	Category        []string `json:"category"`
	PageNo          int      `json:"pageNo"`
	PageSize        int      `json:"pageSize"`
	StartIndex      int      `json:"startIndex"`
	Priority        string   `json:"priority"`
	Search          string   `json:"search"`
	SortBy          string   `json:"sortBy"`
	SortType        string   `json:"sortType"`
	Status          string   `json:"status"`
	UsernamePembuat string   `json:"usernamePembuat"`
	StartDate       string   `json:"startDate" binding:"required"`
	EndDate         string   `json:"endDate" binding:"required"`
}

type CreateTicketRequest struct {
	AssignedTo        string                `json:"assignedTo" form:"assignedTo"`
	Attachment1       *multipart.FileHeader `json:"attachment1" form:"attachment1"`
	Attachment2       *multipart.FileHeader `json:"attachment2" form:"attachment2"`
	Email             string                `json:"email" form:"email"`
	EmailNotification string                `json:"emailNotification" form:"emailNotification"`
	Isi               string                `json:"isi" form:"isi"`
	Judul             string                `json:"judul" form:"judul"`
	Category          string                `json:"category" form:"category"`
	Lokasi            string                `json:"lokasi" form:"lokasi"`
	Prioritas         string                `json:"prioritas" form:"prioritas"`
	Status            string                `json:"status" form:"status"`
	AreaCode          string                `json:"area_code" form:"areaCode"`
	Regional          string                `json:"regional" form:"regional"`
	GrapariId         string                `json:"grapari_id" form:"grapariId"`
	TerminalId        string                `json:"terminalId" form:"terminalId"`
	UserPembuat       string                `json:"userPembuat" form:"userPembuat"`
	TicketCode        string                `json:"ticketCode" form:"ticketCode"`
}

type UpdateTicketRequest struct {
	AssignedTo       string    `json:"assignedTo" form:"assignedTo"`
	Email            string    `json:"email" form:"email"`
	Judul            string    `json:"judul" form:"judul"`
	Category         string    `json:"category" form:"category"`
	Lokasi           string    `json:"lokasi" form:"lokasi"`
	Prioritas        string    `json:"prioritas" form:"prioritas"`
	Status           string    `json:"status" form:"status"`
	TerminalId       string    `json:"terminalId" form:"terminalId"`
	TicketCode       string    `json:"ticketCode" form:"ticketCode"`
	TotalWaktu       string    `json:"totalWaktu" form:"totalWaktu"`
	UsernamePembalas string    `json:"usernamePembalas"`
	TglDiperbarui    time.Time `json:"tglDiperbarui"`
}

type ReplyTicket struct {
	TicketCode        string                `json:"ticketCode" form:"ticketCode"`
	UsernamePengirim  string                `json:"usernamePengirim" form:"usernamePengirim"`
	EmailNotification string                `json:"emailNotification" form:"emailNotification"`
	Status            string                `json:"status" form:"status"`
	Isi               string                `json:"isi" form:"isi"`
	Attachment1       *multipart.FileHeader `json:"attachment1" form:"attachment1"`
	Attachment2       *multipart.FileHeader `json:"attachment2" form:"attachment2"`
	TglDibuat         time.Time             `json:"tglDibuat" form:"tglDibuat"`
}

type UpdateTicketStatusRequest struct {
	Status        string    `json:"status" form:"status"`
	TicketCode    string    `json:"ticketCode" form:"ticketCode"`
	TglDiperbarui time.Time `json:"tglDiperbarui"`
}
