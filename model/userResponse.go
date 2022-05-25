package model

import "svc-myg-ticketing/entity"

type GetUserResponse struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	// Password   string        `json:"password"`
	Phone      string        `json:"phone"`
	Status     string        `json:"status"`
	Area       string        `json:"area"`
	Regional   string        `json:"regional"`
	CreatedAt  string        `json:"createdAt"`
	UpdatedAt  string        `json:"updatedAt"`
	Roles      []entity.Role `json:"roles" gorm:"foreignKey:Id"`
	TerminalId string        `json:"terminalId"`
	RuleId     int           `json:"ruleId"`
	GrapariId  string        `json:"grapariId"`
}

type UserOptions struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type GetUserGroupByRoleResponse struct {
	Id      int           `json:"id" gorm:"primaryKey"`
	Label   string        `json:"label"`
	Options []UserOptions `json:"options"`
}

type GetUserGroupByRole struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Label   string `json:"label"`
	Options string `json:"options"`
}
