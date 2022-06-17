package repository

import (
	"svc-myg-ticketing/entity"
)

type EmailNotifRepositoryInterface interface {
	CreateEmailNotif(request entity.EmailNotif) error
	GetEmailNotif() ([]entity.EmailNotif, error)
}

func (repo *repository) CreateEmailNotif(request entity.EmailNotif) error {

	error := repo.db.Table("email_notif").Create(&request).Error

	return error
}

func (repo *repository) GetEmailNotif() ([]entity.EmailNotif, error) {
	var email_notif []entity.EmailNotif

	error := repo.db.Raw("SELECT * FROM email_notif ORDER BY email ASC").Find(&email_notif).Error

	return email_notif, error
}
