package model

type GetGrapariRequest struct {
	Regional  []string `json:"regional"`
	AreaCode  []string `json:"areaCode"`
	GrapariId []string `json:"grapariId"`
	Status    string   `json:"status"`
}
