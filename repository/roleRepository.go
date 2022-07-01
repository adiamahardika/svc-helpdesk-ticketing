package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type RoleRepositoryInterface interface {
	GetRole() ([]entity.Role, error)
	CreateRole(request model.CreateRoleRequest) ([]entity.Role, error)
	UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error)
	DeleteRole(Id int) error
	GetDetailRole(request model.GetRoleRequest) ([]entity.Role, error)
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT roles.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_permission.id, 'name', ticketing_permission.name, 'code', ticketing_permission.permission_code)) AS list_permission FROM ticketing_role_has_permission INNER JOIN roles ON (roles.id = ticketing_role_has_permission.id_role) INNER JOIN ticketing_permission ON (ticketing_role_has_permission.id_permission = ticketing_permission.id) GROUP BY ticketing_role_has_permission.id_role, roles.id ORDER BY roles.name ASC").Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request model.CreateRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("INSERT INTO roles(name) VALUES(?) RETURNING roles.*", request.Name).Find(&role).Error

	return role, error
}

func (repo *repository) UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error) {
	var role []model.GetRoleResponse

	error := repo.db.Raw("UPDATE roles SET name = @Name WHERE id = @Id RETURNING roles.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(Id int) error {
	var role entity.Role

	error := repo.db.Raw("UPDATE roles SET is_active = ? WHERE id = ? RETURNING roles.*", "false", Id).Find(&role).Error

	return error
}

func (repo *repository) GetDetailRole(request model.GetRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT roles.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_permission.id, 'name', ticketing_permission.name, 'code', ticketing_permission.permission_code)) AS list_permission FROM ticketing_role_has_permission INNER JOIN roles ON (roles.id = ticketing_role_has_permission.id_role) INNER JOIN ticketing_permission ON (ticketing_role_has_permission.id_permission = ticketing_permission.id) WHERE roles.id = @Id GROUP BY ticketing_role_has_permission.id_role, roles.id ORDER BY roles.name ASC", request).Find(&role).Error

	return role, error
}
