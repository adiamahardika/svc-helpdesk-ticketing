package model

import "svc-myg-ticketing/entity"

type CreateRoleRequest struct {
	Name           string              `json:"name"`
	ListPermission []entity.Permission `json:"listPermission" gorm:"foreignKey:Id"`
}
