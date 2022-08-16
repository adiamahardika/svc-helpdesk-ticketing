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

var grapariRepository = &repositoryMock.RepositoryMock{Mock: mock.Mock{}}
var grapariService = service.GrapariService(grapariRepository)

func Test_Service_Grapari_Get(t *testing.T) {

	date := time.Now()
	tests := []struct {
		name           string
		request        *model.GetGrapariRequest
		expectedReturn []entity.MsGrapari
		expectedError  error
	}{
		{
			name: "Success",
			request: &model.GetGrapariRequest{
				Regional:  []string{},
				AreaCode:  []string{},
				GrapariId: []string{},
				Status:    "",
			},
			expectedReturn: []entity.MsGrapari{
				{
					Id:        3,
					GrapariId: "GRP210",
					Name:      "GraPARI Aimas",
					Regional:  "Puma",
					Area:      "4",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        2,
					GrapariId: "GRP001",
					Name:      "GraPARI Alia",
					Regional:  "Central",
					Area:      "2",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        82,
					GrapariId: "GRP080",
					Name:      "GraPARI Ambassador",
					Regional:  "Central",
					Area:      "2",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		grapariRepository.Mock.On("GetGrapari", test.request).Return(test.expectedReturn, test.expectedError)

		t.Run(test.name, func(t *testing.T) {
			result, error := grapariService.GetGrapari(test.request)
			require.Equal(t, test.expectedReturn, result)
			require.Equal(t, test.expectedError, error)
		})
	}

}

func Benchmark_Service_Grapari_Get(b *testing.B) {

	date := time.Now()
	benchmarks := []struct {
		name           string
		request        *model.GetGrapariRequest
		expectedReturn []entity.MsGrapari
	}{
		{
			name: "Success",
			request: &model.GetGrapariRequest{
				Regional:  []string{},
				AreaCode:  []string{},
				GrapariId: []string{},
				Status:    "",
			},
			expectedReturn: []entity.MsGrapari{
				{
					Id:        3,
					GrapariId: "GRP210",
					Name:      "GraPARI Aimas",
					Regional:  "Puma",
					Area:      "4",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        2,
					GrapariId: "GRP001",
					Name:      "GraPARI Alia",
					Regional:  "Central",
					Area:      "2",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
				{
					Id:        82,
					GrapariId: "GRP080",
					Name:      "GraPARI Ambassador",
					Regional:  "Central",
					Area:      "2",
					Status:    "Active",
					CreatedAt: date,
					UpdatedAt: date,
				},
			},
		},
	}

	for _, benchmark := range benchmarks {
		areaRepository.Mock.On("GetGrapari", benchmark.request).Return(benchmark.expectedReturn, nil)

		for index := 0; index < b.N; index++ {
			b.Run(benchmark.name, func(b *testing.B) {
				grapariService.GetGrapari(benchmark.request)
			})
		}
	}
}
