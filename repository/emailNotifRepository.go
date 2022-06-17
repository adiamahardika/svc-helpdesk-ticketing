package repository

import (
	"svc-myg-ticketing/entity"
)

type EmailNotifRepositoryInterface interface {
	CreateEmailNotif(request entity.EmailNotif) error
	GetEmailNotif() ([]entity.EmailNotif, error)
	UpdateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error)
	DeleteEmailNotif(id int) error
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

func (repo *repository) UpdateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error) {
	var email_notif entity.EmailNotif

	error := repo.db.Raw("UPDATE email_notif SET email = @Email, updated_at = @UpdatedAt WHERE id = @Id RETURNING email_notif.*", request).Find(&email_notif).Error

	return email_notif, error
}

func (repo *repository) DeleteEmailNotif(id int) error {

	var email_notif entity.EmailNotif

	error := repo.db.Raw("DELETE FROM email_notif WHERE id = ? RETURNING email_notif.*", id).Find(&email_notif).Error

	return error
}
