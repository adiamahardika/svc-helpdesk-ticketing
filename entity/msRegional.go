package entity

import "time"

type MsRegional struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Regional  string    `json:"areaCode"`
	Area      string    `json:"area"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
