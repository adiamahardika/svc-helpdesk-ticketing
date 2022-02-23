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
	Name             string    `json:"name"`
	Parent           string    `json:"parent"`
	UpdateAt         time.Time `json:"updateAt"`
	AdditionalInput1 string    `json:"additionalInput1"`
	AdditionalInput2 string    `json:"additionalInput2"`
	AdditionalInput3 string    `json:"additionalInput3"`
}
