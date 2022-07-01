package repository

import (
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

	error := repo.db.Raw("SELECT users.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_role.id, 'name', ticketing_role.name)) AS roles FROM users LEFT OUTER JOIN ticketing_user_has_role ON (users.id = ticketing_user_has_role.id_user) LEFT OUTER JOIN ticketing_role ON (ticketing_role.id = ticketing_user_has_role.id_role) WHERE users.name LIKE @Search OR users.username LIKE @Search OR users.email LIKE @Search GROUP BY ticketing_user_has_role.id_user, users.id ORDER BY users.name ASC LIMIT @Size OFFSET @StartIndex", model.GetUserRequest{
		Search:     "%" + request.Search + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) CountUser(request model.GetUserRequest) (int, error) {
	var total_data int

	error := repo.db.Raw("SELECT COUNT(*) as total_data FROM users LEFT OUTER JOIN ticketing_user_has_role ON (users.id = ticketing_user_has_role.id_user) LEFT OUTER JOIN ticketing_role ON (ticketing_role.id = ticketing_user_has_role.id_role) WHERE users.name LIKE @Search OR users.username LIKE @Search OR users.email LIKE @Search", model.GetUserRequest{
		Search: "%" + request.Search + "%",
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) GetUserDetail(request string) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("SELECT users.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_role.id, 'name', ticketing_role.name)) AS roles, JSON_AGG(DISTINCT user_has_area.area) AS area_code, JSON_AGG(DISTINCT user_has_region.regional) AS regional, JSON_AGG(DISTINCT user_has_grapari.grapari_id) AS grapari_id FROM users LEFT OUTER JOIN ticketing_user_has_role ON (users.id = ticketing_user_has_role.id_user) LEFT OUTER JOIN ticketing_role ON (ticketing_role.id = ticketing_user_has_role.id_role) LEFT OUTER JOIN user_has_area ON (users.username = user_has_area.username) LEFT OUTER JOIN user_has_region ON (users.username = user_has_region.username) LEFT OUTER JOIN user_has_grapari ON (users.username = user_has_grapari.username) WHERE users.username = ? GROUP BY ticketing_user_has_role.id_user, users.id ORDER BY users.name ASC", request).Find(&user).Error

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

	error := repo.db.Raw("SELECT ticketing_role.id, ticketing_role.name AS label, JSON_AGG(JSON_BUILD_OBJECT('label', users.name, 'value', users.username)) AS options FROM ticketing_user_has_role INNER JOIN ticketing_role ON (ticketing_role.id = ticketing_user_has_role.id_role) INNER JOIN users ON (users.id = ticketing_user_has_role.id_user) WHERE role.is_active = 'true' AND users.status = 'Active'  GROUP BY ticketing_role.name, ticketing_role.id ORDER BY ticketing_role.name ASC").Find(&user).Error

	return user, error
}
