package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
	"time"
)

type EmailNotifServiceInterface interface {
	CreateEmailNotif(request entity.EmailNotif) (entity.EmailNotif, error)
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
