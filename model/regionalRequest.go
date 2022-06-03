package model

type GetRegionalRequest struct {
	Regional []string `json:"regional"`
	AreaCode []string `json:"areaCode"`
	Status   string   `json:"status"`
}
