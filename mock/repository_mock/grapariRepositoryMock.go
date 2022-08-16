package repositorymock_test

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

func (repo *RepositoryMock) GetGrapari(request *model.GetGrapariRequest) ([]entity.MsGrapari, error) {

	arguments := repo.Mock.Called(request)

	grapari := arguments.Get(0).([]entity.MsGrapari)

	return grapari, nil
}
