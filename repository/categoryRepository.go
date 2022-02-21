package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type CategoryRepositoryInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]entity.Category, error)
	CreateCategory(request entity.Category) (entity.Category, error)
}

func (repo *repository) GetCategory(request model.GetCategoryRequest) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category WHERE is_active LIKE @IsActive ORDER BY name", model.GetCategoryRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&category).Error

	return category, error
}

func (repo *repository) CreateCategory(request entity.Category) (entity.Category, error) {
	var category entity.Category

	error := repo.db.Table("category").Create(&request).Error

	return category, error
}
