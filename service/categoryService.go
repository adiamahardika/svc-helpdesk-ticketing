package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"
)

type CategoryServiceInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]entity.Category, error)
	CreateCategory(request model.CreateCategoryRequest) (entity.Category, error)
	UpdateCategory(request entity.Category) (entity.Category, error)
	DeleteCategory(Id int) error
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

func (categoryService *categoryService) CreateCategory(request model.CreateCategoryRequest) (entity.Category, error) {
	date_now := time.Now()

	category_request := entity.Category{
		Name:              request.Name,
		CodeLevel:         request.CodeLevel,
		Parent:            request.Parent,
		AdditionalInput_1: request.AdditionalInput1,
		AdditionalInput_2: request.AdditionalInput2,
		AdditionalInput_3: request.AdditionalInput3,
		IsActive:          "true",
		UpdateAt:          date_now,
	}

	_, error := categoryService.repository.CreateCategory(category_request)

	return category_request, error
}

func (categoryService *categoryService) UpdateCategory(request entity.Category) (entity.Category, error) {
	date_now := time.Now()

	request.UpdateAt = date_now

	category, error := categoryService.repository.UpdateCategory(request)

	return category, error
}

func (categoryService *categoryService) DeleteCategory(Id int) error {

	error := categoryService.repository.DeleteCategory(Id)

	return error
}
