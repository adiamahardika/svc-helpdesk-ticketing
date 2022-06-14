package entity

import "time"

type SubCategory struct {
	Id         int       `json:"id" gorm:"primaryKey" gorm:"->"`
	Name       string    `json:"name"`
	IdCategory int       `json:"idCategory"`
	Priority   string    `json:"priority"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
