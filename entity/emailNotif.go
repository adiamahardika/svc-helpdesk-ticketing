package entity

import "time"

type EmailNotif struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
