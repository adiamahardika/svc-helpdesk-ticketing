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

var categoryRepository = &repositoryMock.RepositoryMock{Mock: mock.Mock{}}
var subCategoryRepository = &repositoryMock.RepositoryMock{Mock: mock.Mock{}}
var categoryService = service.CategoryService(categoryRepository, subCategoryRepository)

func TestGetCategoryService(t *testing.T) {
	date, _ := time.Parse("0001-01-01T00:00:00Z", time.RFC1123Z)
	tests := []struct {
		name              string
		request           *model.GetCategoryRequest
		expectedReturn    []model.CreateCategoryRequest
		expectedTotalPage float64
		expectedError     error
	}{{
		name: "Success Get Category",
		request: &model.GetCategoryRequest{
			Size:     100,
			PageNo:   0,
			IsActive: "true",
		},
		expectedReturn: []model.CreateCategoryRequest{
			{
				Id:   80,
				Name: "Aplikasi",
				SubCategory: []entity.SubCategory{{
					Id:         66,
					Name:       "Tidak bisa login",
					IdCategory: 80,
					Priority:   "High",
					CreatedAt:  date,
					UpdatedAt:  date,
				},
					{
						Id:         67,
						Name:       "Error/hang",
						IdCategory: 80,
						Priority:   "Critical",
						CreatedAt:  date,
						UpdatedAt:  date,
					},
					{
						Id:         68,
						Name:       "Gagal update",
						IdCategory: 80,
						Priority:   "High",
						CreatedAt:  date,
						UpdatedAt:  date,
					}},
				IsActive: "true",
				UpdateAt: date,
			},
			{
				Id:   84,
				Name: "Bayar Halo / Isi Pulsa Via Qris",
				SubCategory: []entity.SubCategory{{
					Id:         51,
					Name:       "Saldo debit/ cc terpotong dan transaksi gagal",
					IdCategory: 84,
					Priority:   "High",
					CreatedAt:  date,
					UpdatedAt:  date,
				}},
				IsActive: "true",
				UpdateAt: date,
			},
			{
				Id:   78,
				Name: "Bill Acceptor Rupiah",
				SubCategory: []entity.SubCategory{{
					Id:         40,
					Name:       "Tidak dapat menerima uang/uang tereject",
					IdCategory: 78,
					Priority:   "High",
					CreatedAt:  date,
					UpdatedAt:  date,
				},
					{
						Id:         41,
						Name:       "Mati/error",
						IdCategory: 78,
						Priority:   "Critical",
						CreatedAt:  date,
						UpdatedAt:  date,
					},
					{
						Id:         42,
						Name:       "Uang tersangkut di dalam Bill acceptor",
						IdCategory: 78,
						Priority:   "High",
						CreatedAt:  date,
						UpdatedAt:  date,
					}},
				IsActive: "true",
				UpdateAt: date,
			},
		},
		expectedTotalPage: 1,
		expectedError:     nil,
	}}

	for _, test := range tests {
		categoryRepository.Mock.On("CountCategory", test.request).Return(int(test.expectedTotalPage), test.expectedError)
		categoryRepository.Mock.On("GetCategory", test.request).Return([]entity.Category{
			{
				Id:          80,
				Name:        "Aplikasi",
				SubCategory: `[{"id":66,"name":"Tidak bisa login","idCategory":80,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":67,"name":"Error/hang","idCategory":80,"priority":"Critical","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":68,"name":"Gagal update","idCategory":80,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]`,
				IsActive:    "true",
				UpdateAt:    date,
			},
			{
				Id:          84,
				Name:        "Bayar Halo / Isi Pulsa Via Qris",
				SubCategory: `[{"id":51,"name":"Saldo debit/ cc terpotong dan transaksi gagal","idCategory":84,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]`,
				IsActive:    "true",
				UpdateAt:    date,
			},
			{
				Id:          78,
				Name:        "Bill Acceptor Rupiah",
				SubCategory: `[{"id":40,"name":"Tidak dapat menerima uang/uang tereject","idCategory":78,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":41,"name":"Mati/error","idCategory":78,"priority":"Critical","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":42,"name":"Uang tersangkut di dalam Bill acceptor","idCategory":78,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]`,
				IsActive:    "true",
				UpdateAt:    date,
			},
		}, nil)

		t.Run(test.name, func(t *testing.T) {
			response, total_page, error := categoryService.GetCategory(test.request)
			require.Equal(t, test.expectedReturn, response)
			require.Equal(t, test.expectedTotalPage, total_page)
			require.Equal(t, test.expectedError, error)
		})
	}
}
