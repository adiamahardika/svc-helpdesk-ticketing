package model

import "time"

type ReportResponse struct {
	Id                int    `json:"id" gorm:"primaryKey"`
	Judul             string `json:"judul"`
	UsernamePembuat   string `json:"usernamePembuat"`
	UpdatedBy         string `json:"updatedBy" form:"updatedBy"`
	Prioritas         string `json:"prioritas"`
	TglDibuat         string `json:"tglDibuat"`
	TglDiperbarui     string `json:"tglDiperbarui"`
	TotalWaktu        string `json:"totalWaktu"`
	Status            string `json:"status"`
	TicketCode        string `json:"ticketCode"`
	Category          string `json:"category"`
	Email             string `json:"email"`
	AssignedTo        string `json:"assignedTo"`
	EmailNotification string `json:"emailNotification"`
	Isi               string `json:"isi" gorm:"->"`
	AreaCode          string `json:"areaCode"`
	AreaName          string `json:"areaName" gorm:"->"`
	Regional          string `json:"regional"`
	GrapariId         string `json:"grapariId"`
	GrapariName       string `json:"grapariName" gorm:"->"`
	TerminalId        string `json:"terminalId"`
	Lokasi            string `json:"lokasi"`
	CategoryName      string `json:"categoryName" gorm:"->"`
	UserPembuat       string `json:"userPembuat" gorm:"->"`
	Assignee          string `json:"assignee" gorm:"->"`
	SubCategory       string `json:"subCategory"`
	StartTime         string `json:"startTime" gorm:"->"`
	StartBy           string `json:"startBy" gorm:"->"`
	CloseTime         string `json:"closeTime" gorm:"->"`
	CloseBy           string `json:"closeBy" gorm:"->"`
	AssigningTime     string `json:"assigningTime"`
	AssigningBy       string `json:"assigningBy"`
	VisitStatus       string `json:"visitStatus" gorm:"->"`
}

type GetCountReportByStatusResponse struct {
	Date    time.Time `json:"date"`
	New     int       `json:"new"`
	Process int       `json:"process"`
	Finish  int       `json:"finish"`
}
