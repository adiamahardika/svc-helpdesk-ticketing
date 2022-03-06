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

type UpdateUserRequest struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Area       string    `json:"area"`
	Regional   string    `json:"regional"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Role       string    `json:"role"`
	TerminalId string    `json:"terminalId"`
	RuleId     int       `json:"ruleId"`
	GrapariId  string    `json:"grapariId"`
}

type ChangePassRequest struct {
	Username    string    `json:"username"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResetPassword struct {
	Username    string    `json:"username"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}
