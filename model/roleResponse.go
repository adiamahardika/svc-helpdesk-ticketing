package model

import "svc-myg-ticketing/entity"

type GetRoleResponse struct {
	Id             int                 `json:"id" gorm:"primaryKey"`
	Name           string              `json:"name"`
	ListPermission []entity.Permission `json:"list_permission" gorm:"foreignKey:Id"`
}
