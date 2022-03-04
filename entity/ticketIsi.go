package entity

type TicketIsi struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	KodeTicket       string `json:"kodeTicket"`
	UsernamePengirim string `json:"usernamePengirim"`
	Isi              string `json:"isi"`
	Attachment1      string `json:"attachment1"`
	Attachment2      string `json:"attachment2"`
	TglDibuat        string `json:"tglDibuat"`
	UrlAttachment1   string `json:"urlAttachment1"`
	UrlAttachment2   string `json:"urlAttachment2"`
}
