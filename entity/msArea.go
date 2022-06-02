package entity

import "time"

type MsArea struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	AreaCode  string    `json:"areaCode"`
	AreaName  string    `json:"areaName"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
