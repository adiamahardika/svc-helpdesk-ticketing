package repository

import "svc-myg-ticketing/entity"

type RoleRepositoryInterface interface {
	GetRole() ([]entity.Role, error)
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT role.*, JSON_AGG(JSON_BUILD_OBJECT('id', permission.id, 'name', permission.name, 'code', permission.permission_code)) AS list_permission FROM role_has_permission INNER JOIN role ON (role.id = role_has_permission.id_role) inner JOIN permission ON (role_has_permission.id_permission = permission.id) GROUP by role_has_permission.id_role, role.id ORDER BY role.name ASC").Find(&role).Error

	return role, error
}
