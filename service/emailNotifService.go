package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
	"time"
)

type EmailNotifServiceInterface interface {
	CreateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error)
	GetEmailNotif() ([]entity.EmailNotif, error)
	UpdateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error)
	DeleteEmailNotif(id int) error
}

type emailNotifService struct {
	emailNotifRepository repository.EmailNotifRepositoryInterface
}

func EmailNotifService(emailNotifRepository repository.EmailNotifRepositoryInterface) *emailNotifService {
	return &emailNotifService{emailNotifRepository}
}

func (emailNotifService *emailNotifService) CreateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error) {
	date_now := time.Now()

	request.CreatedAt = date_now
	request.UpdatedAt = date_now
	error := emailNotifService.emailNotifRepository.CreateEmailNotif(request)

	return request, error
}

func (emailNotifService *emailNotifService) GetEmailNotif() ([]entity.EmailNotif, error) {

	email_notif, error := emailNotifService.emailNotifRepository.GetEmailNotif()

	return email_notif, error
}

func (emailNotifService *emailNotifService) UpdateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error) {
	date_now := time.Now()

	request.UpdatedAt = date_now
	email_notif, error := emailNotifService.emailNotifRepository.UpdateEmailNotif(request)

	return email_notif, error
}

func (emailNotifService *emailNotifService) DeleteEmailNotif(id int) error {

	error := emailNotifService.emailNotifRepository.DeleteEmailNotif(id)

	return error
}
