package entity

import "time"

type Category struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	CodeLevel   string    `json:"codeLevel"`
	Name        string    `json:"name"`
	Parent      string    `json:"parent"`
	SubCategory string    `json:"subCategory"`
	IsActive    string    `json:"isActive"`
	UpdateAt    time.Time `json:"updateAt"`
}
