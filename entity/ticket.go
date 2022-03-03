package entity

type Ticket struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	Judul            string `json:"judul"`
	UsernamePembuat  string `json:"usernamePembuat"`
	UsernamePembalas string `json:"usernamePembalas"`
	Prioritas        string `json:"prioritas"`
	TglDibuat        string `json:"tglDibuat"`
	TglDiperbarui    string `json:"tglDiperbarui"`
	TotalWaktu       string `json:"totalWaktu"`
	Status           string `json:"status"`
	KodeTicket       string `json:"kodeTicket"`
	Kategori         string `json:"kategori"`
	Lokasi           string `json:"lokasi"`
	TerminalId       string `json:"terminalId"`
	Email            string `json:"email"`
	AssignedTo       string `json:"assignedTo"`
}
