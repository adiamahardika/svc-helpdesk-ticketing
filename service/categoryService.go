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
	GetCategory(request model.GetCategoryRequest) ([]model.CreateCategoryRequest, float64, error)
	CreateCategory(request model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	UpdateCategory(request model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	DeleteCategory(Id int) error
	GetDetailCategory(request string) ([]model.CreateCategoryRequest, error)
}

type categoryService struct {
	categoryRepository    repository.CategoryRepositoryInterface
	subCategoryRepository repository.SubCategoryRepositoryInterface
}

func CategoryService(categoryRepository repository.CategoryRepositoryInterface, subCategoryRepository repository.SubCategoryRepositoryInterface) *categoryService {
	return &categoryService{categoryRepository, subCategoryRepository}
}

func (categoryService *categoryService) GetCategory(request model.GetCategoryRequest) ([]model.CreateCategoryRequest, float64, error) {
	var response []model.CreateCategoryRequest

	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size
	total_data, error := categoryService.categoryRepository.CountCategory(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.Size))

	category, error := categoryService.categoryRepository.GetCategory(request)

	for _, value := range category {
		var sub_category []entity.SubCategory
		json.Unmarshal([]byte(value.SubCategory), &sub_category)

		response = append(response, model.CreateCategoryRequest{
			Id:          value.Id,
			Name:        value.Name,
			SubCategory: sub_category,
			IsActive:    value.IsActive,
			UpdateAt:    value.UpdateAt,
		})
	}

	return response, total_pages, error
}

func (categoryService *categoryService) CreateCategory(request model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {
	var sub_category []entity.SubCategory
	var response model.CreateCategoryRequest
	date_now := time.Now()

	request.UpdateAt = date_now
	request.IsActive = "true"

	category, error := categoryService.categoryRepository.CreateCategory(request)

	if error == nil {
		for _, value := range request.SubCategory {
			sub_category = append(sub_category, entity.SubCategory{
				Name:       value.Name,
				IdCategory: category.Id,
				Priority:   value.Priority,
				CreatedAt:  date_now,
				UpdatedAt:  date_now,
			})
		}
		sub_category, error = categoryService.subCategoryRepository.CreateSubCategory(sub_category)

		response = model.CreateCategoryRequest{
			Id:          category.Id,
			Name:        category.Name,
			SubCategory: sub_category,
			IsActive:    category.IsActive,
			UpdateAt:    category.UpdateAt,
		}
	}

	return response, error
}

func (categoryService *categoryService) UpdateCategory(request model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {
	var sub_category []entity.SubCategory
	date_now := time.Now()

	request.UpdateAt = date_now
	category, error := categoryService.categoryRepository.UpdateCategory(request)

	if error == nil {
		error = categoryService.subCategoryRepository.DeleteSubCategory(request.Id)

		if error == nil {
			for _, value := range request.SubCategory {
				sub_category = append(sub_category, entity.SubCategory{
					Name:       value.Name,
					IdCategory: request.Id,
					Priority:   value.Priority,
					CreatedAt:  date_now,
					UpdatedAt:  date_now,
				})
			}

			sub_category, error = categoryService.subCategoryRepository.CreateSubCategory(sub_category)
			category.SubCategory = sub_category
		}
	}

	return category, error
}

func (categoryService *categoryService) DeleteCategory(Id int) error {

	error := categoryService.categoryRepository.DeleteCategory(Id)

	return error
}

func (categoryService *categoryService) GetDetailCategory(request string) ([]model.CreateCategoryRequest, error) {
	var response []model.CreateCategoryRequest

	category, error := categoryService.categoryRepository.GetDetailCategory(request)

	for _, value := range category {
		var sub_category []entity.SubCategory
		json.Unmarshal([]byte(value.SubCategory), &sub_category)

		response = append(response, model.CreateCategoryRequest{
			Id:          value.Id,
			Name:        value.Name,
			SubCategory: sub_category,
			IsActive:    value.IsActive,
			UpdateAt:    value.UpdateAt,
		})
	}

	return response, error
}
