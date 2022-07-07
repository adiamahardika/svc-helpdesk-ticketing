package model

import (
	"svc-myg-ticketing/entity"
	"time"
)

type GetCategoryRequest struct {
	Size       int    `json:"size"`
	PageNo     int    `json:"pageNo"`
	StartIndex int    `json:"startIndex"`
	SortBy     string `json:"sortBy"`
	OrderBy    string `json:"orderBy"`
	IsActive   string `json:"isActive"`
}

type CreateCategoryRequest struct {
	Id          int                   `json:"id" gorm:"primaryKey"`
	Name        string                `json:"name"`
	SubCategory []*entity.SubCategory `json:"subCategory" gorm:"foreignKey:Id"`
	IsActive    string                `json:"isActive"`
	UpdateAt    time.Time             `json:"updateAt"`
}
