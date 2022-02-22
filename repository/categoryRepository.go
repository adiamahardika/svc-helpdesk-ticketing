package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type CategoryRepositoryInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]entity.Category, error)
	CreateCategory(request entity.Category) (entity.Category, error)
	UpdateCategory(request entity.Category) (entity.Category, error)
	DeleteCategory(Id int) error
	GetDetailCategory(request string) ([]entity.Category, error)
}

func (repo *repository) GetCategory(request model.GetCategoryRequest) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category WHERE is_active LIKE @IsActive ORDER BY name ASC LIMIT @Size OFFSET @StartIndex", model.GetCategoryRequest{
		IsActive:   "%" + request.IsActive + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&category).Error

	return category, error
}

func (repo *repository) CreateCategory(request entity.Category) (entity.Category, error) {
	var category entity.Category

	error := repo.db.Table("category").Create(&request).Error

	return category, error
}

func (repo *repository) UpdateCategory(request entity.Category) (entity.Category, error) {

	var category entity.Category

	error := repo.db.Raw("UPDATE category SET name = @Name, code_level = @CodeLevel, parent = @Parent, additional_input_1 = @AdditionalInput_1, additional_input_2 = @AdditionalInput_2, additional_input_3 = @AdditionalInput_3, update_at = @UpdateAt WHERE id = @Id RETURNING category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) DeleteCategory(Id int) error {
	var category entity.Category

	error := repo.db.Raw("UPDATE category SET is_active = ? WHERE id = ? RETURNING category.*", "false", Id).Find(&category).Error

	return error
}

func (repo *repository) GetDetailCategory(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category WHERE code_level = ?", request).Find(&category).Error

	return category, error
}