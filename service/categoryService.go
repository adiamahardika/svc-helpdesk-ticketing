package service

import (
	"math"
	"strings"
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
	GetDetailCategory(request string) ([]entity.Category, []entity.Category, []entity.Category, []entity.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepositoryInterface
}

func CategoryService(repository repository.CategoryRepositoryInterface) *categoryService {
	return &categoryService{repository}
}

func (categoryService *categoryService) GetCategory(request model.GetCategoryRequest) ([]entity.Category, error) {

	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size

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

func (categoryService *categoryService) GetDetailCategory(request string) ([]entity.Category, []entity.Category, []entity.Category, []entity.Category, error) {

	category, error := categoryService.repository.GetDetailCategory(request)
	var parent_1 []entity.Category
	var parent_2 []entity.Category
	var parent_3 []entity.Category

	split := strings.Split(category[0].Parent, ".")

	if len(split) == 2 {
		parent_1, error = categoryService.repository.GetDetailCategory(category[0].Parent)
		parent_2 = nil
		parent_3 = nil
	} else if len(split) == 3 {
		parent_2, error = categoryService.repository.GetDetailCategory(category[0].Parent)
		parent_1, error = categoryService.repository.GetDetailCategory(parent_1[0].Parent)
		parent_3 = nil
	} else if len(split) == 4 {
		parent_3, error = categoryService.repository.GetDetailCategory(category[0].Parent)
		parent_2, error = categoryService.repository.GetDetailCategory(parent_3[0].Parent)
		parent_1, error = categoryService.repository.GetDetailCategory(parent_2[0].Parent)
	} else {
		parent_1 = nil
		parent_2 = nil
		parent_3 = nil
	}

	return category, parent_1, parent_2, parent_3, error
}
