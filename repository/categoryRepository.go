package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type CategoryRepositoryInterface interface {
	GetCategory(request *model.GetCategoryRequest) ([]*entity.Category, error)
	CountCategory(request *model.GetCategoryRequest) (int, error)
	CreateCategory(request *model.CreateCategoryRequest) (*model.CreateCategoryRequest, error)
	UpdateCategory(request *model.CreateCategoryRequest) (*model.CreateCategoryRequest, error)
	DeleteCategory(Id *int) error
	GetDetailCategory(request *string) ([]*entity.Category, error)
	GetCategoryByParentDesc(request string) ([]entity.Category, error)
	GetCategoryByParentAsc(request string) ([]entity.Category, error)
}

func (repo *repository) GetCategory(request *model.GetCategoryRequest) ([]*entity.Category, error) {
	var category []*entity.Category

	error := repo.db.Raw("SELECT ticketing_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_sub_category.id, 'idCategory', ticketing_sub_category.id_category, 'name', ticketing_sub_category.name, 'priority', ticketing_sub_category.priority)) AS sub_category FROM ticketing_category LEFT OUTER JOIN ticketing_sub_category ON (ticketing_category.id = ticketing_sub_category.id_category) WHERE is_active LIKE @IsActive GROUP BY ticketing_category.id ORDER BY name ASC LIMIT @Size OFFSET @StartIndex", model.GetCategoryRequest{
		IsActive:   "%" + request.IsActive + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&category).Error

	return category, error
}

func (repo *repository) CountCategory(request *model.GetCategoryRequest) (int, error) {
	var total_data int

	error := repo.db.Raw("SELECT COUNT(*) as total_data FROM ticketing_category WHERE is_active LIKE @IsActive", model.GetCategoryRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) CreateCategory(request *model.CreateCategoryRequest) (*model.CreateCategoryRequest, error) {
	var category *model.CreateCategoryRequest

	error := repo.db.Raw("INSERT INTO ticketing_category(name, is_active, update_at) VALUES(@Name, @IsActive, @UpdateAt) RETURNING ticketing_category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) UpdateCategory(request *model.CreateCategoryRequest) (*model.CreateCategoryRequest, error) {

	var category *model.CreateCategoryRequest

	error := repo.db.Raw("UPDATE ticketing_category SET name = @Name, update_at = @UpdateAt WHERE id = @Id RETURNING ticketing_category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) DeleteCategory(Id *int) error {
	var category *entity.Category

	error := repo.db.Raw("UPDATE ticketing_category SET is_active = ? WHERE id = ? RETURNING ticketing_category.*", "false", Id).Find(&category).Error

	return error
}

func (repo *repository) GetDetailCategory(request *string) ([]*entity.Category, error) {
	var category []*entity.Category

	error := repo.db.Raw("SELECT ticketing_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_sub_category.id, 'idCategory', ticketing_sub_category.id_category, 'name', ticketing_sub_category.name, 'priority', ticketing_sub_category.priority)) AS sub_category FROM ticketing_category LEFT OUTER JOIN ticketing_sub_category ON (ticketing_category.id = ticketing_sub_category.id_category) WHERE ticketing_category.id = ? GROUP BY ticketing_category.id", request).Find(&category).Error

	return category, error
}

func (repo *repository) GetCategoryByParentDesc(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM ticketing_category WHERE parent = ? ORDER BY update_at DESC", request).Find(&category).Error

	return category, error
}

func (repo *repository) GetCategoryByParentAsc(request string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM ticketing_category WHERE parent = ? ORDER BY update_at ASC", request).Find(&category).Error

	return category, error
}
