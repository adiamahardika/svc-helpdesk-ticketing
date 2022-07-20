package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type GrapariServiceInterface interface {
	GetGrapari(request *model.GetGrapariRequest) ([]entity.MsGrapari, error)
}

type grapariService struct {
	repository repository.GrapariRepositotyInterface
}

func GrapariService(repository repository.GrapariRepositotyInterface) *grapariService {
	return &grapariService{repository}
}

func (grapariService *grapariService) GetGrapari(request *model.GetGrapariRequest) ([]entity.MsGrapari, error) {

	grapari, error := grapariService.repository.GetGrapari(request)

	return grapari, error
}
