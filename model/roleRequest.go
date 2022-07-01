package model

import "svc-myg-ticketing/entity"

type CreateRoleRequest struct {
	Name           string              `json:"name"`
	ListPermission []entity.Permission `json:"listPermission" gorm:"foreignKey:Id"`
}

type UpdateRoleRequest struct {
	Id             int                 `json:"id" gorm:"primaryKey"`
	Name           string              `json:"name"`
	ListPermission []entity.Permission `json:"listPermission" gorm:"foreignKey:Id"`
}

type GetRoleRequest struct {
	Id int `json:"id"`
}
