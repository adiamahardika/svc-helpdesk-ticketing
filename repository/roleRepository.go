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

	error := repo.db.Raw("SELECT ticketing_role.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_permission.id, 'name', ticketing_permission.name, 'code', ticketing_permission.permission_code)) AS list_permission FROM ticketing_role_has_permission INNER JOIN ticketing_role ON (ticketing_role.id = ticketing_role_has_permission.id_role) INNER JOIN ticketing_permission ON (ticketing_role_has_permission.id_permission = ticketing_permission.id) WHERE ticketing_role.is_active = 'true' GROUP BY ticketing_role_has_permission.id_role, ticketing_role.id ORDER BY ticketing_role.name ASC").Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request model.CreateRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("INSERT INTO ticketing_role(name) VALUES(?) RETURNING ticketing_role.*", request.Name).Find(&role).Error

	return role, error
}

func (repo *repository) UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error) {
	var role []model.GetRoleResponse

	error := repo.db.Raw("UPDATE ticketing_role SET name = @Name WHERE id = @Id RETURNING ticketing_role.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(Id int) error {
	var role entity.Role

	error := repo.db.Raw("UPDATE ticketing_role SET is_active = ? WHERE id = ? RETURNING ticketing_role.*", "false", Id).Find(&role).Error

	return error
}

func (repo *repository) GetDetailRole(request model.GetRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT ticketing_role.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_permission.id, 'name', ticketing_permission.name, 'code', ticketing_permission.permission_code)) AS list_permission FROM ticketing_role_has_permission INNER JOIN ticketing_role ON (ticketing_role.id = ticketing_role_has_permission.id_role) INNER JOIN ticketing_permission ON (ticketing_role_has_permission.id_permission = ticketing_permission.id) INNER JOIN ticketing_user_has_role ON (ticketing_user_has_role.id_role = ticketing_role.id) WHERE ticketing_user_has_role.id_user = @IdUser GROUP BY ticketing_role_has_permission.id_role, ticketing_role.id ORDER BY ticketing_role.name ASC", request).Find(&role).Error

	return role, error
}
