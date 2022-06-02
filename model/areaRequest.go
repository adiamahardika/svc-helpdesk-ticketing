package model

type GetAreaRequest struct {
	AreaCode string `json:"areaCode"`
	AreaName string `json:"areaName"`
	Status   string `json:"status"`
}
