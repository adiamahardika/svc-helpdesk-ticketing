package service_test

import (
	"svc-myg-ticketing/entity"
	repositoryMock "svc-myg-ticketing/mock/repository_mock"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var areaRepository = &repositoryMock.RepositoryMock{Mock: mock.Mock{}}
var areaService = service.AreaService(areaRepository)

func Test_Service_Area_Get(t *testing.T) {

	date := time.Now()
	tests := []struct {
		name           string
		request        *model.GetAreaRequest
		expectedReturn []entity.MsArea
		expectedError  error
	}{
		{
			name: "Success",
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
			require.Equal(t, test.expectedReturn, result)
			require.Equal(t, test.expectedError, error)
		})
	}

}

func Benchmark_Service_Area_Get(b *testing.B) {

	date := time.Now()
	benchmarks := []struct {
		name           string
		request        *model.GetAreaRequest
		expectedReturn []entity.MsArea
	}{
		{
			name: "Get Area",
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
		},
	}

	for _, benchmark := range benchmarks {
		areaRepository.Mock.On("GetArea", benchmark.request).Return(benchmark.expectedReturn, nil)

		for index := 0; index < b.N; index++ {
			b.Run(benchmark.name, func(b *testing.B) {
				areaService.GetArea(benchmark.request)
			})
		}
	}
}
