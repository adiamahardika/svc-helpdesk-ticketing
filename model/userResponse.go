package model

import "svc-myg-ticketing/entity"

type GetUserResponse struct {
	Id         int           `json:"id" gorm:"primaryKey"`
	Username   string        `json:"username"`
	Name       string        `json:"name"`
	Password   string        `json:"password"`
	Area       string        `json:"area"`
	Regional   string        `json:"regional"`
	CreatedAt  string        `json:"createdAt"`
	UpdatedAt  string        `json:"updatedAt"`
	Roles      []entity.Role `json:"roles" gorm:"foreignKey:Id"`
	TerminalId string        `json:"terminalId"`
	RuleId     int           `json:"ruleId"`
	GrapariId  string        `json:"grapariId"`
}
