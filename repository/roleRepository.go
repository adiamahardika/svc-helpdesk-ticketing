package repository

import "svc-myg-ticketing/entity"

type RoleRepositoryInterface interface {
	GetRole() ([]entity.Role, error)
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("WITH list_permission AS (SELECT * FROM permission LEFT OUTER JOIN role_has_permission ON (permission.id = role_has_permission.id_permission)) SELECT role.* FROM role LEFT OUTER JOIN role_has_permission ON (role.id IN role_has_permission.id_role) ORDER BY role.name ASC").Find(&role).Error

	return role, error
}
