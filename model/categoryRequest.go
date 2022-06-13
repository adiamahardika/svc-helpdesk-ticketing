package model

import "time"

type GetCategoryRequest struct {
	Size       int    `json:"size"`
	PageNo     int    `json:"pageNo"`
	StartIndex int    `json:"startIndex"`
	SortBy     string `json:"sortBy"`
	OrderBy    string `json:"orderBy"`
	IsActive   string `json:"isActive"`
}

type CreateCategoryRequest struct {
	Name     string    `json:"name"`
	UpdateAt time.Time `json:"updateAt"`
}
