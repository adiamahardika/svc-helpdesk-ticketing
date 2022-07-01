package model

import "svc-myg-ticketing/entity"

type GetRoleResponse struct {
	Id             int                 `json:"id" gorm:"primaryKey"`
	Name           string              `json:"name"`
	GuardName      string              `json:"guardName"`
	ListPermission []entity.Permission `json:"listPermission" gorm:"foreignKey:Id"`
}
