package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type CategoryRepositoryInterface interface {
	GetCategory(request model.GetCategoryRequest) ([]entity.Category, error)
	CountCategory(request model.GetCategoryRequest) (int, error)
	CreateCategory(request entity.Category) (entity.Category, error)
	UpdateCategory(request entity.Category) (entity.Category, error)
	DeleteCategory(Id int) error
	GetDetailCategory(request string) ([]entity.Category, error)
	GetCategoryByParentDesc(request string) ([]entity.Category, error)
	GetCategoryByParentAsc(request string) ([]entity.Category, error)
}

func (repo *repository) GetCategory(request model.GetCategoryRequest) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT category.*, JSON_AGG(JSON_BUILD_OBJECT('id', sub_category.id, 'idCategory', sub_category.id_category, 'name', sub_category.name, 'priority', sub_category.priority)) AS sub_category FROM category LEFT OUTER JOIN sub_category ON (category.id = sub_category.id_category) WHERE is_active LIKE @IsActive GROUP BY category.id ORDER BY name ASC LIMIT @Size OFFSET @StartIndex", model.GetCategoryRequest{
		IsActive:   "%" + request.IsActive + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&category).Error

	return category, error
}

func (repo *repository) CountCategory(request model.GetCategoryRequest) (int, error) {
	var total_data int

	error := repo.db.Raw("SELECT COUNT(*) as total_data FROM category WHERE is_active LIKE @IsActive", model.GetCategoryRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) CreateCategory(request entity.Category) (entity.Category, error) {
	var category entity.Category

	error := repo.db.Table("category").Create(&request).Error

	return category, error
}

func (repo *repository) UpdateCategory(request entity.Category) (entity.Category, error) {

	var category entity.Category

	error := repo.db.Raw("UPDATE category SET name = @Name, update_at = @UpdateAt WHERE id = @Id RETURNING category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) DeleteCategory(Id int) error {
	var category entity.Category

	error := repo.db.Raw("UPDATE category SET is_active = ? WHERE id = ? RETURNING category.*", "false", Id).Find(&category).Error

	return error
}

func (repo *repository) GetDetailCategory(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT category.*, JSON_AGG(JSON_BUILD_OBJECT('id', sub_category.id, 'idCategory', sub_category.id_category, 'name', sub_category.name, 'priority', sub_category.priority)) AS sub_category FROM category LEFT OUTER JOIN sub_category ON (category.id = sub_category.id_category) WHERE category.id = ? GROUP BY category.id", request).Find(&category).Error

	return category, error
}

func (repo *repository) GetCategoryByParentDesc(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category WHERE parent = ? ORDER BY update_at DESC", request).Find(&category).Error

	return category, error
}

func (repo *repository) GetCategoryByParentAsc(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category WHERE parent = ? ORDER BY update_at ASC", request).Find(&category).Error

	return category, error
}
