package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type UserRepositoryInterface interface {
	GetUser(request model.GetUserRequest) ([]entity.User, error)
	GetUserDetail(request string) (entity.User, error)
	CountUser(request model.GetUserRequest) (int, error)
	DeleteUser(id int) error
	CreateUser(request model.CreateUserRequest) (entity.User, error)
	CheckUsername(request string) ([]entity.User, error)
	UpdateUser(request model.UpdateUserRequest) (entity.User, error)
	ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error)
	UpdateUserStatus(request model.UpdateUserStatus) (entity.User, error)
	GetUserGroupByRole() ([]model.GetUserGroupByRole, error)
}

func (repo *repository) GetUser(request model.GetUserRequest) ([]entity.User, error) {
	var user []entity.User
	var role string

	if request.Role != 0 {
		role = "WHERE model_has_roles.role_id = @Role"
	}

	query := fmt.Sprintf("SELECT * FROM (SELECT users.*, JSON_AGG(JSON_BUILD_OBJECT('id', roles.id, 'name', roles.name)) AS roles FROM users LEFT OUTER JOIN model_has_roles ON (users.id = model_has_roles.model_id) LEFT OUTER JOIN roles ON (roles.id = model_has_roles.role_id) %s GROUP BY model_has_roles.model_id, users.id ORDER BY users.name ASC) AS tbl WHERE tbl.name LIKE @Search OR tbl.username LIKE @Search OR tbl.email LIKE @Search LIMIT @Size OFFSET @StartIndex", role)

	error := repo.db.Raw(query, model.GetUserRequest{
		Search:     "%" + request.Search + "%",
		Role:       request.Role,
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) CountUser(request model.GetUserRequest) (int, error) {
	var total_data int
	var role string

	if request.Role != 0 {
		role = "WHERE model_has_roles.role_id = @Role"
	}

	query := fmt.Sprintf("SELECT COUNT(*) as total_data FROM (SELECT users.* FROM users LEFT OUTER JOIN model_has_roles ON (users.id = model_has_roles.model_id) LEFT OUTER JOIN roles ON (roles.id = model_has_roles.role_id) %s) AS tbl WHERE tbl.name LIKE @Search OR tbl.username LIKE @Search OR tbl.email LIKE @Search", role)

	error := repo.db.Raw(query, model.GetUserRequest{
		Search: "%" + request.Search + "%",
		Role:   request.Role,
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) GetUserDetail(request string) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("SELECT users.*, model_has_roles.role_id AS rule_id, JSON_AGG(JSON_BUILD_OBJECT('id', roles.id, 'name', roles.name)) AS roles, JSON_AGG(DISTINCT user_has_area.area) AS area_code, JSON_AGG(DISTINCT user_has_region.regional) AS regional, JSON_AGG(DISTINCT user_has_grapari.grapari_id) AS grapari_id FROM users LEFT OUTER JOIN model_has_roles ON (users.id = model_has_roles.model_id) LEFT OUTER JOIN roles ON (roles.id = model_has_roles.role_id) LEFT OUTER JOIN user_has_area ON (users.username = user_has_area.username) LEFT OUTER JOIN user_has_region ON (users.username = user_has_region.username) LEFT OUTER JOIN user_has_grapari ON (users.username = user_has_grapari.username) WHERE users.username = ? GROUP BY users.id, model_has_roles.role_id ORDER BY users.name ASC", request).Find(&user).Error

	return user, error
}

func (repo *repository) DeleteUser(id int) error {
	var users entity.User

	error := repo.db.Raw("UPDATE users SET status = ? WHERE id = ? RETURNING users.*", "Inactive", id).Find(&users).Error

	return error
}

func (repo *repository) CreateUser(request model.CreateUserRequest) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("INSERT INTO users(name,username,password,email,phone,status,updated_at,created_at) VALUES(@Name, @Username, @Password, @Email, @Phone, @Status, @UpdatedAt, @CreatedAt) RETURNING users.*", request).Find(&user).Error

	return user, error
}

func (repo *repository) CheckUsername(request string) ([]entity.User, error) {
	var user []entity.User

	error := repo.db.Raw("SELECT * FROM users WHERE username = @Username", model.CreateUserRequest{
		Username: request,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) UpdateUser(request model.UpdateUserRequest) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("UPDATE users SET name = @Name, email = @Email, phone = @Phone WHERE id = @Id RETURNING users.*", request).Find(&user).Error

	return user, error
}

func (repo *repository) ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error) {

	var user model.GetUserResponse

	error := repo.db.Raw("UPDATE users SET password = @NewPassword, updated_at = @UpdatedAt WHERE username = @Username RETURNING users.*", request).Find(&user).Error

	return user, error
}

func (repo *repository) UpdateUserStatus(request model.UpdateUserStatus) (entity.User, error) {

	var user entity.User

	error := repo.db.Raw("UPDATE users SET status = @Status, updated_at = @UpdatedAt WHERE username = @Username RETURNING users.*", request).Find(&user).Error

	return user, error
}

func (repo *repository) GetUserGroupByRole() ([]model.GetUserGroupByRole, error) {
	var user []model.GetUserGroupByRole

	error := repo.db.Raw("SELECT roles.id, roles.name AS label, JSON_AGG(JSON_BUILD_OBJECT('label', users.name, 'value', users.username)) AS options FROM model_has_roles INNER JOIN roles ON (roles.id = model_has_roles.role_id) INNER JOIN users ON (users.id = model_has_roles.model_id) WHERE users.status = 'Active'  GROUP BY roles.name, roles.id ORDER BY roles.name ASC").Find(&user).Error

	return user, error
}
