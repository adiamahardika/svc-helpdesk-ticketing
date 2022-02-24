package repository

import "svc-myg-ticketing/model"

type RoleHasPermissionRepositoryInterface interface {
	CreateRoleHasPermission(request []model.CreateRoleHasPermissionRequest) error
}

func (repo *repository) CreateRoleHasPermission(request []model.CreateRoleHasPermissionRequest) error {

	error := repo.db.Table("role_has_permission").Create(&request).Error

	return error
}
