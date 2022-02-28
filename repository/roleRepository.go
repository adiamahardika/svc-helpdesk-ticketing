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
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT role.*, JSON_AGG(JSON_BUILD_OBJECT('id', permission.id, 'name', permission.name, 'code', permission.permission_code)) AS list_permission FROM role_has_permission INNER JOIN role ON (role.id = role_has_permission.id_role) INNER JOIN permission ON (role_has_permission.id_permission = permission.id) WHERE role.is_active = 'true' GROUP BY role_has_permission.id_role, role.id ORDER BY role.name ASC").Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request model.CreateRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("INSERT INTO role(name) VALUES(?) RETURNING role.*", request.Name).Find(&role).Error

	return role, error
}

func (repo *repository) UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error) {
	var role []model.GetRoleResponse

	error := repo.db.Raw("UPDATE role SET name = @Name WHERE id = @Id RETURNING role.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(Id int) error {
	var role entity.Role

	error := repo.db.Raw("UPDATE role SET is_active = ? WHERE id = ? RETURNING role.*", "false", Id).Find(&role).Error

	return error
}
