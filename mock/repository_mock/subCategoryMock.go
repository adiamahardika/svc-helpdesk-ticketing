package repositorymock_test

import (
	"svc-myg-ticketing/entity"
)

func (repo *RepositoryMock) CreateSubCategory(request []*entity.SubCategory) ([]entity.SubCategory, error) {

	arguments := repo.Mock.Called(request)

	sub_category := arguments.Get(0).([]entity.SubCategory)

	return sub_category, nil
}

func (repo *RepositoryMock) DeleteSubCategory(id_category *int) error {

	return nil
}

func (repo *RepositoryMock) GetSubCategory() ([]entity.SubCategory, error) {

	arguments := repo.Mock.Called()

	sub_category := arguments.Get(0).([]entity.SubCategory)

	return sub_category, nil
}
