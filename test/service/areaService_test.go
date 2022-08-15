package service_test

import (
	"svc-myg-ticketing/entity"
	repositorymock_test "svc-myg-ticketing/mock/repository_mock"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var areaRepository = &repositorymock_test.RepositoryMock{Mock: mock.Mock{}}
var areaService = service.AreaService(areaRepository)

func TestGetAreaService(t *testing.T) {

	date := time.Now()
	tests := []struct {
		name           string
		request        *model.GetAreaRequest
		expectedReturn []entity.MsArea
		expectedError  error
	}{
		{
			name: "Success Get Area",
			request: &model.GetAreaRequest{
				AreaCode: []string{""},
				AreaName: "",
				Status:   "",
			},
			expectedReturn: []entity.MsArea{
				{
					Id:        3,
					AreaCode:  "1",
					AreaName:  "Sumatra",
					Status:    "A",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        2,
					AreaCode:  "2",
					AreaName:  "Jabotabek-Jabar",
					Status:    "A",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        1,
					AreaCode:  "3",
					AreaName:  "Jawa-Balinusra",
					Status:    "A",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        4,
					AreaCode:  "4",
					AreaName:  "Pamasuka",
					Status:    "A",
					CreatedAt: date,
					UpdatedAt: date,
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		areaRepository.Mock.On("GetArea", test.request).Return(test.expectedReturn, test.expectedError)

		t.Run(test.name, func(t *testing.T) {
			result, error := areaService.GetArea(test.request)
			assert.Equal(t, test.expectedReturn, result)
			assert.Equal(t, test.expectedError, error)
		})
	}

}
