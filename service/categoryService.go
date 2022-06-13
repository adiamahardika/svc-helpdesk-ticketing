package service

import (
	"encoding/json"
	"math"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"
)

type CategoryServiceInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]model.GetCategoryResponse, float64, error)
	CreateCategory(request model.CreateCategoryRequest) (entity.Category, error)
	UpdateCategory(request entity.Category) (entity.Category, error)
	DeleteCategory(Id int) error
	GetDetailCategory(request string) ([]entity.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepositoryInterface
}

func CategoryService(repository repository.CategoryRepositoryInterface) *categoryService {
	return &categoryService{repository}
}

func (categoryService *categoryService) GetCategory(request model.GetCategoryRequest) ([]model.GetCategoryResponse, float64, error) {
	var response []model.GetCategoryResponse

	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size
	total_data, error := categoryService.repository.CountCategory(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.Size))

	category, error := categoryService.repository.GetCategory(request)

	for _, value := range category {
		var sub_category []entity.SubCategory
		json.Unmarshal([]byte(value.SubCategory), &sub_category)

		response = append(response, model.GetCategoryResponse{
			Id:          value.Id,
			Name:        value.Name,
			SubCategory: sub_category,
			IsActive:    value.IsActive,
			UpdateAt:    value.UpdateAt,
		})
	}

	return response, total_pages, error
}

func (categoryService *categoryService) CreateCategory(request model.CreateCategoryRequest) (entity.Category, error) {
	date_now := time.Now()

	category_request := entity.Category{
		Name:     request.Name,
		IsActive: "true",
		UpdateAt: date_now,
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

func (categoryService *categoryService) GetDetailCategory(request string) ([]entity.Category, error) {

	category, error := categoryService.repository.GetDetailCategory(request)

	return category, error
}
