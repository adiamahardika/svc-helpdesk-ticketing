package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type RegionalServiceInterface interface {
	GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error)
}

type regionalService struct {
	repository repository.RegionalRepositoryInterface
}

func RegionalService(repository repository.RegionalRepositoryInterface) *regionalService {
	return &regionalService{repository}
}

func (regionalService *regionalService) GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error) {

	regional, error := regionalService.repository.GetRegional(request)

	return regional, error
}
