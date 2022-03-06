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
	ChangePassword(request model.ChangePassRequest) (entity.User, error)
}

func (repo *repository) GetUser(request model.GetUserRequest) ([]entity.User, error) {
	var user []entity.User

	error := repo.db.Raw("SELECT users.*, JSON_AGG(JSON_BUILD_OBJECT('id', role.id, 'name', role.name)) AS roles FROM users LEFT OUTER JOIN user_has_role ON (users.id = user_has_role.id_user) LEFT OUTER JOIN role ON (role.id = user_has_role.id_role) WHERE users.name LIKE @Search OR users.username LIKE @Search OR users.email LIKE @Search GROUP BY user_has_role.id_user, users.id ORDER BY users.name ASC LIMIT @Size OFFSET @StartIndex", model.GetUserRequest{
		Search:     "%" + request.Search + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) CountUser(request model.GetUserRequest) (int, error) {
	var total_data int

	error := repo.db.Raw("SELECT COUNT(*) as total_data FROM users LEFT OUTER JOIN user_has_role ON (users.id = user_has_role.id_user) LEFT OUTER JOIN role ON (role.id = user_has_role.id_role) WHERE users.name LIKE @Search OR users.username LIKE @Search OR users.email LIKE @Search", model.GetUserRequest{
		Search: "%" + request.Search + "%",
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) GetUserDetail(request string) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("SELECT users.*, JSON_AGG(JSON_BUILD_OBJECT('id', role.id, 'name', role.name)) AS roles FROM users LEFT OUTER JOIN user_has_role ON (users.id = user_has_role.id_user) LEFT OUTER JOIN role ON (role.id = user_has_role.id_role) WHERE users.username = ? GROUP BY user_has_role.id_user, users.id ORDER BY users.name ASC", request).Find(&user).Error

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

	error := repo.db.Raw("UPDATE users SET name = @Name, email = @Email, phone = @Phone, area = @Area, regional = @Regional, updated_at = @UpdatedAt, terminal_id = @TerminalId, rule_id = @RuleId, grapari_id = @GrapariId WHERE id = @Id RETURNING users.*", request).Find(&user).Error

	return user, error
}

func (repo *repository) ChangePassword(request model.ChangePassRequest) (entity.User, error) {

	var user entity.User

	error := repo.db.Raw("UPDATE users SET password = @NewPassword, updated_at = @UpdatedAt WHERE username = @Username RETURNING users.*", request).Find(&user).Error

	return user, error
}
