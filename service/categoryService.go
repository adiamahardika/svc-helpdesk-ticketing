package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type CategoryServiceInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]entity.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepositoryInterface
}

func CategoryService(repository repository.CategoryRepositoryInterface) *categoryService {
	return &categoryService{repository}
}

func (categoryService *categoryService) GetCategory(request model.GetCategoryRequest) ([]entity.Category, error) {

	return categoryService.repository.GetCategory(request)
}
