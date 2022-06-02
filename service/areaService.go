package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type AreaServiceInterface interface {
	GetArea(request model.GetAreaRequest) ([]entity.MsArea, error)
}

type areaService struct {
	repository repository.AreaRepositoryInterface
}

func AreaService(repository repository.AreaRepositoryInterface) *areaService {
	return &areaService{repository}
}

func (areaService *areaService) GetArea(request model.GetAreaRequest) ([]entity.MsArea, error) {

	area, error := areaService.repository.GetArea(request)

	return area, error
}
