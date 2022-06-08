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
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	Id        int       `json:"id" gorm:"primaryKey" binding:"required"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updatedAt"`
	Role      string    `json:"role"`
}

type ChangePassRequest struct {
	Username    string    `json:"username" binding:"required"`
	OldPassword string    `json:"oldPassword" binding:"required"`
	NewPassword string    `json:"newPassword" binding:"required"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ResetPassword struct {
	Username    string    `json:"username" binding:"required"`
	NewPassword string    `json:"newPassword" binding:"required"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateUserStatus struct {
	Username  string    `json:"username" binding:"required"`
	Status    string    `json:"status" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt"`
}
