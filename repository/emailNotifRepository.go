package repository

import (
	"svc-myg-ticketing/entity"
)

type EmailNotifRepositoryInterface interface {
	CreateEmailNotif(request entity.EmailNotif) error
}

func (repo *repository) CreateEmailNotif(request entity.EmailNotif) error {

	error := repo.db.Table("email_notif").Create(&request).Error

	return error
}
