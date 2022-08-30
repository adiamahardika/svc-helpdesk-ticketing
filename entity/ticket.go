package entity

import "time"

type Ticket struct {
	Id                int       `json:"id" gorm:"primaryKey"`
	Judul             string    `json:"judul"`
	UsernamePembuat   string    `json:"usernamePembuat"`
	UpdatedBy         string    `json:"updatedBy" form:"updatedBy"`
	Prioritas         string    `json:"prioritas"`
	TglDibuat         time.Time `json:"tglDibuat"`
	TglDiperbarui     time.Time `json:"tglDiperbarui"`
	TotalWaktu        string    `json:"totalWaktu"`
	Status            string    `json:"status"`
	TicketCode        string    `json:"ticketCode"`
	Category          string    `json:"category"`
	Email             string    `json:"email"`
	AssignedTo        string    `json:"assignedTo"`
	EmailNotification string    `json:"emailNotification"`
	Isi               string    `json:"isi" gorm:"->"`
	AreaCode          string    `json:"areaCode"`
	AreaName          string    `json:"areaName" gorm:"->"`
	Regional          string    `json:"regional"`
	GrapariId         string    `json:"grapariId"`
	GrapariName       string    `json:"grapariName" gorm:"->"`
	TerminalId        string    `json:"terminalId"`
	Lokasi            string    `json:"lokasi"`
	CategoryName      string    `json:"categoryName" gorm:"->"`
	UserPembuat       string    `json:"userPembuat" gorm:"->"`
	Assignee          string    `json:"assignee" gorm:"->"`
	SubCategory       string    `json:"subCategory"`
	StartTime         time.Time `json:"startTime" gorm:"->"`
	StartBy           string    `json:"startBy" gorm:"->"`
	CloseTime         time.Time `json:"closeTime" gorm:"->"`
	CloseBy           string    `json:"closeBy" gorm:"->"`
	AssigningTime     time.Time `json:"assigningTime"`
	AssigningBy       string    `json:"assigningBy"`
	VisitStatus       string    `json:"visitStatus" gorm:"->"`
}
