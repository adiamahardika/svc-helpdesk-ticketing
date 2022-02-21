package model

type GetCategoryRequest struct {
	Size     string `json:"size"`
	PageNo   string `json:"page_no"`
	SortBy   string `json:"sort_by"`
	OrderBy  string `json:"order_by"`
	IsActive string `json:"is_active"`
}
