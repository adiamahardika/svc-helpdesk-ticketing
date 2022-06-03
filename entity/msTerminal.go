package entity

import "time"

type MsTerminal struct {
	Id               int       `json:"id" gorm:"primaryKey"`
	TerminalId       string    `json:"terminalId"`
	TerminalName     string    `json:"terminalName"`
	GrapariId        string    `json:"grapariId"`
	Regional         string    `json:"regional"`
	Area             string    `json:"area"`
	CtpType          string    `json:"ctpType"`
	TerminalLocation string    `json:"terminalLocation"`
	Kecamatan        string    `json:"kecamatan"`
	KotaKabupaten    string    `json:"kotaKabupaten"`
	KodePos          string    `json:"kodePos"`
	KontakPic        string    `json:"kontakPic"`
	Status           string    `json:"status"`
	CreatedDate      time.Time `json:"createdDate"`
}
