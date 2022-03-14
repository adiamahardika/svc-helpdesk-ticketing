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
	TerminalId        string                `json:"terminalId" form:"terminalId"`
	UserPembuat       string                `json:"userPembuat" form:"userPembuat"`
	TotalWaktu        string                `json:"totalWaktu" form:"totalWaktu"`
	TicketCode        string                `json:"ticketCode" form:"ticketCode"`
}

type UpdateTicketRequest struct {
	AssignedTo    string    `json:"assignedTo" form:"assignedTo"`
	Email         string    `json:"email" form:"email"`
	Judul         string    `json:"judul" form:"judul"`
	Category      string    `json:"category" form:"category"`
	Lokasi        string    `json:"lokasi" form:"lokasi"`
	Prioritas     string    `json:"prioritas" form:"prioritas"`
	Status        string    `json:"status" form:"status"`
	TerminalId    string    `json:"terminalId" form:"terminalId"`
	TicketCode    string    `json:"ticketCode" form:"ticketCode"`
	TglDiperbarui time.Time `json:"tglDiperbarui"`
}
