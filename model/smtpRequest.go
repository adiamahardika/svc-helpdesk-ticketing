package model

type SmtpRequest struct {
	Judul           string `json:"judul"`
	Prioritas       string `json:"prioritas"`
	UsernamePembuat string `json:"usernamePembuat"`
	Status          string `json:"status"`
	TicketCode      string `json:"ticketCode"`
	AreaName        string `json:"areaName" gorm:"->"`
	Regional        string `json:"regional"`
	GrapariName     string `json:"grapariName" gorm:"->"`
	TerminalId      string `json:"terminalId"`
	Lokasi          string `json:"lokasi"`
	Email           string `json:"email"`
	Isi             string `json:"isi"`
	Type            string `json:"type"`
	CategoryName    string `json:"categoryName" gorm:"->"`
	SubCategory     string `json:"subCategory"`
	UserPembuat     string `json:"userPembuat" gorm:"->"`
	Assignee        string `json:"assignee" gorm:"->"`
}
