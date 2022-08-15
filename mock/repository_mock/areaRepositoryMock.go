package repositorymock_test

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

func (repo *RepositoryMock) GetArea(request *model.GetAreaRequest) ([]entity.MsArea, error) {

	arguments := repo.Mock.Called(request)

	area := arguments.Get(0).([]entity.MsArea)

	return area, nil
}
