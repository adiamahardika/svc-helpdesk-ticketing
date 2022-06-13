package model

import (
	"svc-myg-ticketing/entity"
	"time"
)

type GetCategoryResponse struct {
	Id          int                  `json:"id" gorm:"primaryKey"`
	CodeLevel   string               `json:"codeLevel"`
	Name        string               `json:"name"`
	Parent      string               `json:"parent"`
	SubCategory []entity.SubCategory `json:"sub_category"`
	IsActive    string               `json:"isActive"`
	UpdateAt    time.Time            `json:"updateAt"`
}
