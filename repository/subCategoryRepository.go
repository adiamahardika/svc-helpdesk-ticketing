package repository

import "svc-myg-ticketing/entity"

type SubCategoryRepositoryInterface interface {
	CreateSubCategory(request []*entity.SubCategory) ([]entity.SubCategory, error)
	DeleteSubCategory(id_category *int) error
	GetSubCategory() ([]entity.SubCategory, error)
}

func (repo *repository) CreateSubCategory(request []*entity.SubCategory) ([]entity.SubCategory, error) {
	var sub_category []entity.SubCategory

	error := repo.db.Table("ticketing_sub_category").Create(&request).Find(&sub_category).Error

	return sub_category, error
}

func (repo *repository) DeleteSubCategory(id_category *int) error {
	var sub_category *entity.SubCategory

	error := repo.db.Raw("DELETE FROM ticketing_sub_category WHERE id_category = ? RETURNING ticketing_sub_category.*", id_category).Find(&sub_category).Error

	return error
}

func (repo *repository) GetSubCategory() ([]entity.SubCategory, error) {
	var sub_category []entity.SubCategory

	error := repo.db.Raw("SELECT * FROM ticketing_sub_category ORDER BY name ASC").Find(&sub_category).Error

	return sub_category, error
}
