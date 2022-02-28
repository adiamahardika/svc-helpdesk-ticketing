package repository

import "svc-myg-ticketing/entity"

type PermissionRepositoryInterface interface {
	GetPermission() ([]entity.Permission, error)
}

func (repo *repository) GetPermission() ([]entity.Permission, error) {
	var permission []entity.Permission

	error := repo.db.Raw("SELECT * FROM permission ORDER BY name ASC").Find(&permission).Error

	return permission, error
}
