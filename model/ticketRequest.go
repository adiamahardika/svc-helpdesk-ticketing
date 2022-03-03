package model

type GetTicketRequest struct {
	AssignedTo      string   `json:"assignedTo"`
	Category        []string `json:"category"`
	PageNo          int      `json:"pageNo"`
	PageSize        int      `json:"pageSize"`
	StartIndex      int      `json:"startIndex"`
	Priority        string   `json:"priority"`
	Search          string   `json:"search"`
	SortBy          string   `json:"sortBy"`
	SortType        string   `json:"sortType"`
	Status          string   `json:"status"`
	UsernamePembuat string   `json:"usernamePembuat"`
}
