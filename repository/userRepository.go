package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type UserRepositoryInterface interface {
	GetUser(request model.GetUserRequest) ([]entity.User, error)
	GetUserDetail(request string) (entity.User, error)
	CountUser(request model.GetUserRequest) (int, error)
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
