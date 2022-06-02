package entity

import "time"

type MsGrapari struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	GrapariId string    `json:"grapariId"`
	Name      string    `json:"name"`
	Regional  string    `json:"regional"`
	Area      string    `json:"area"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
