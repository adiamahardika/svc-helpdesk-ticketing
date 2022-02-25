package model

import "svc-myg-ticketing/entity"

type GetRoleResponse struct {
	Id             int                 `json:"id" gorm:"primaryKey"`
	Name           string              `json:"name"`
	IsActive       string              `json:"isActive"`
	ListPermission []entity.Permission `json:"listPermission" gorm:"foreignKey:Id"`
}
