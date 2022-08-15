package repositorymock_test

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

func (repo *RepositoryMock) GetCategory(request *model.GetCategoryRequest) ([]entity.Category, error) {

	arguments := repo.Mock.Called(request)

	category := arguments.Get(0).([]entity.Category)

	return category, nil
}

func (repo *RepositoryMock) CountCategory(request *model.GetCategoryRequest) (int, error) {

	arguments := repo.Mock.Called(request)

	category := arguments.Get(0).(int)

	return category, nil
}

func (repo *RepositoryMock) CreateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {

	arguments := repo.Mock.Called(request)

	category := arguments.Get(0).(model.CreateCategoryRequest)

	return category, nil
}

func (repo *RepositoryMock) UpdateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {

	arguments := repo.Mock.Called(request)

	category := arguments.Get(0).(model.CreateCategoryRequest)

	return category, nil
}

func (repo *RepositoryMock) DeleteCategory(id *int) error {

	return nil
}

func (repo *RepositoryMock) GetDetailCategory(request *string) ([]entity.Category, error) {

	arguments := repo.Mock.Called(request)

	category := arguments.Get(0).([]entity.Category)

	return category, nil
}
