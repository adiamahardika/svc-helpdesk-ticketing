package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type RoleHasPermissionRepositoryInterface interface {
	CreateRoleHasPermission(request []model.CreateRoleHasPermissionRequest) error
	DeleteRoleHasPermission(id_role int) error
}

func (repo *repository) CreateRoleHasPermission(request []model.CreateRoleHasPermissionRequest) error {

	error := repo.db.Table("role_has_permission").Create(&request).Error

	return error
}

func (repo *repository) DeleteRoleHasPermission(id_role int) error {

	var role_has_permission entity.RoleHasPermission

	error := repo.db.Raw("DELETE FROM role_has_permission WHERE id_role = ? RETURNING role_has_permission.*", id_role).Find(&role_has_permission).Error

	return error
}
