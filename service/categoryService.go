package service

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
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
	GetDetailCategory(request string) ([]entity.Category, []entity.Category, []entity.Category, []entity.Category, error)
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
			CodeLevel:   value.CodeLevel,
			Parent:      value.Parent,
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
		Name:      request.Name,
		CodeLevel: "",
		Parent:    request.Parent,
		IsActive:  "true",
		UpdateAt:  date_now,
	}

	find_by_parent, error := categoryService.repository.GetCategoryByParentDesc(request.Parent)

	if len(find_by_parent) > 0 {
		last_code_level := find_by_parent[0].CodeLevel
		split := strings.Split(last_code_level, ".")
		last_code := split[len(split)-1]
		var parse_code int
		parse_code, error = strconv.Atoi(last_code)
		parse_code++

		category_request.CodeLevel = request.Parent + "." + strconv.Itoa(parse_code)
	} else {
		category_request.CodeLevel = request.Parent + ".1"
	}

	_, error = categoryService.repository.CreateCategory(category_request)

	return category_request, error
}

func (categoryService *categoryService) UpdateCategory(request entity.Category) (entity.Category, error) {
	date_now := time.Now()

	find_by_parent, error := categoryService.repository.GetCategoryByParentDesc(request.Parent)
	find_by_code_level, error := categoryService.repository.GetCategoryByParentAsc(request.CodeLevel)

	if len(find_by_parent) > 0 {
		last_code_level := find_by_parent[0].CodeLevel
		split := strings.Split(last_code_level, ".")
		last_code := split[len(split)-1]
		var parse_code int
		parse_code, error = strconv.Atoi(last_code)

		request.CodeLevel = request.Parent + "." + strconv.Itoa(parse_code+1)
	} else {
		request.CodeLevel = request.Parent + ".1"
	}
	request.UpdateAt = date_now
	category, error := categoryService.repository.UpdateCategory(request)

	if len(find_by_code_level) > 0 {
		for index, value := range find_by_code_level {
			number := index + 1
			parse_index := strconv.Itoa(number)
			value.CodeLevel = request.CodeLevel + "." + parse_index
			value.Parent = request.CodeLevel
			value.UpdateAt = date_now
			_, error = categoryService.repository.UpdateCategory(value)
		}
	}

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
		parent_1, error = categoryService.repository.GetDetailCategory(parent_2[0].Parent)
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
