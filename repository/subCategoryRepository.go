package repository

import "svc-myg-ticketing/entity"

type SubCategoryRepositoryInterface interface {
	CreateSubCategory(request []entity.SubCategory) ([]entity.SubCategory, error)
	DeleteSubCategory(id_category int) error
}

func (repo *repository) CreateSubCategory(request []entity.SubCategory) ([]entity.SubCategory, error) {
	var sub_category []entity.SubCategory

	error := repo.db.Table("sub_category").Create(&request).Find(&sub_category).Error

	return sub_category, error
}

func (repo *repository) DeleteSubCategory(id_category int) error {
	var sub_category entity.SubCategory

	error := repo.db.Raw("DELETE FROM sub_category WHERE id_category = ? RETURNING sub_category.*", id_category).Find(&sub_category).Error

	return error
}
