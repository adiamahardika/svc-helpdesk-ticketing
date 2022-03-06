package model

import "time"

type GetUserRequest struct {
	Search     string `json:"search"`
	Size       int    `json:"size"`
	PageNo     int    `json:"pageNo"`
	StartIndex int    `json:"startIndex"`
}

type CreateUserRequest struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
