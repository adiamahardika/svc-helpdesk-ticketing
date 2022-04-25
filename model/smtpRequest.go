package model

type SmtpRequest struct {
	Judul           string `json:"judul"`
	Prioritas       string `json:"prioritas"`
	UsernamePembuat string `json:"usernamePembuat"`
	Author          string `json:"author"`
	Status          string `json:"status"`
	TicketCode      string `json:"ticketCode"`
	Lokasi          string `json:"lokasi"`
	TerminalId      string `json:"terminalId"`
	Email           string `json:"email"`
	Isi             string `json:"isi"`
	Type            string `json:"type"`
}
