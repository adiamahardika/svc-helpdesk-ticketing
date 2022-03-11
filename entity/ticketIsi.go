package entity

import "time"

type TicketIsi struct {
	Id               int       `json:"id" gorm:"primaryKey"`
	TicketCode       string    `json:"ticketCode"`
	UsernamePengirim string    `json:"usernamePengirim"`
	Isi              string    `json:"isi"`
	Attachment1      string    `json:"attachment1"`
	Attachment2      string    `json:"attachment2"`
	TglDibuat        time.Time `json:"tglDibuat"`
	UrlAttachment1   string    `json:"urlAttachment1"`
	UrlAttachment2   string    `json:"urlAttachment2"`
}
