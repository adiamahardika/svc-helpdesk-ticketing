package model

import (
	"svc-myg-ticketing/entity"
	"time"
)

type GetCategoryResponse struct {
	Id          int                  `json:"id" gorm:"primaryKey"`
	Name        string               `json:"name"`
	SubCategory []entity.SubCategory `json:"sub_category"`
	IsActive    string               `json:"isActive"`
	UpdateAt    time.Time            `json:"updateAt"`
}
