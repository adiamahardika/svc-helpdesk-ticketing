package model

import "time"

type GetCategoryRequest struct {
	Size     string `json:"size"`
	PageNo   string `json:"page_no"`
	SortBy   string `json:"sort_by"`
	OrderBy  string `json:"order_by"`
	IsActive string `json:"is_active"`
}

type CreateCategoryRequest struct {
	Name             string    `json:"name"`
	CodeLevel        string    `json:"code_level"`
	Parent           string    `json:"parent"`
	UpdateAt         time.Time `json:"update_at"`
	AdditionalInput1 string    `json:"additional_input_1"`
	AdditionalInput2 string    `json:"additional_input_2"`
	AdditionalInput3 string    `json:"additional_input_3"`
}
