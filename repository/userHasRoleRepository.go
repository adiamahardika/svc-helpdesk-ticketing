package repository

import "svc-myg-ticketing/entity"

type UserHasRoleRepositoryInterface interface {
	CreateUserHasRole(id_user int, id_role int) error
	DeleteUserHasRole(id_user int) error
}

func (repo *repository) CreateUserHasRole(id_user int, id_role int) error {

	var user_has_role entity.UserHasRole

	error := repo.db.Raw("INSERT INTO ticketing_user_has_role(id_user, id_role) VALUES(?, ?) RETURNING ticketing_user_has_role.*", id_user, id_role).Find(&user_has_role).Error

	return error
}

func (repo *repository) DeleteUserHasRole(id_user int) error {

	var user_has_role entity.UserHasRole

	error := repo.db.Raw("DELETE FROM ticketing_user_has_role WHERE id_user = ? RETURNING ticketing_user_has_role.*", id_user).Find(&user_has_role).Error

	return error
}
