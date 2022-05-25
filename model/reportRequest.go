package model

type GetReportRequest struct {
	AssignedTo      string   `json:"assignedTo"`
	Category        []string `json:"category"`
	Priority        []string `json:"priority"`
	Status          []string `json:"status"`
	UsernamePembuat []string `json:"usernamePembuat"`
	StartDate       string   `json:"startDate" binding:"required"`
	EndDate         string   `json:"endDate" binding:"required"`
}
