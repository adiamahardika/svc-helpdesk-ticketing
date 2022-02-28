package model

type GetUserRequest struct {
	Search     string `json:"search"`
	Size       int    `json:"size"`
	PageNo     int    `json:"pageNo"`
	StartIndex int    `json:"startIndex"`
}
