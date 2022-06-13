package entity

import "time"

type Category struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	SubCategory string    `json:"subCategory" gorm:"->"`
	IsActive    string    `json:"isActive"`
	UpdateAt    time.Time `json:"updateAt"`
}
