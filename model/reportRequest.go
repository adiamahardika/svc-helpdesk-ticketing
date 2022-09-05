package model

type GetReportRequest struct {
	AssignedTo      string   `json:"assignedTo"`
	Category        []string `json:"category"`
	AreaCode        []string `json:"areaCode"`
	Regional        []string `json:"regional"`
	GrapariId       []string `json:"grapariId"`
	Priority        []string `json:"priority"`
	Status          []string `json:"status"`
	UsernamePembuat []string `json:"usernamePembuat"`
	StartDate       string   `json:"startDate" binding:"required"`
	EndDate         string   `json:"endDate" binding:"required"`
}

type GetCountReportByStatusRequest struct {
	AreaCode  []string `json:"areaCode"`
	Regional  []string `json:"regional"`
	GrapariId []string `json:"grapariId"`
	StartDate string   `json:"startDate" binding:"required"`
	EndDate   string   `json:"endDate" binding:"required"`
}
