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

func TestCreateCategoryService(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedReturn model.CreateCategoryRequest
		expectedError  error
	}{{
		name: "Success Create Category",
		request: &model.CreateCategoryRequest{
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Name:     "Mati/error",
					Priority: "Critical",
				},
				{
					Name:     "Uang tersangkut di dalam Bill acceptor",
					Priority: "High",
				},
			},
			UpdateAt: date,
		},
		expectedReturn: model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:         41,
					Name:       "Mati/error",
					IdCategory: 78,
					Priority:   "Critical",
					UpdatedAt:  date,
				},
				{
					Id:         42,
					Name:       "Uang tersangkut di dalam Bill acceptor",
					IdCategory: 78,
					Priority:   "High",
					UpdatedAt:  date,
				},
			},
			IsActive: "true",
			UpdateAt: date,
		},
		expectedError: nil,
	}}

	for _, test := range tests {
		var request_sc []*entity.SubCategory
		for _, value := range test.expectedReturn.SubCategory {
			request_sc = append(request_sc, &entity.SubCategory{
				Name:       value.Name,
				IdCategory: test.expectedReturn.Id,
				Priority:   value.Priority,
				CreatedAt:  date,
				UpdatedAt:  date,
			})
		}

		subCategoryRepository.Mock.On("CreateSubCategory", request_sc).Return(test.expectedReturn.SubCategory, test.expectedError)
		categoryRepository.Mock.On("CreateCategory", test.request).Return(test.expectedReturn, nil)

		t.Run(test.name, func(t *testing.T) {
			response, error := categoryService.CreateCategory(test.request)
			require.Equal(t, test.expectedReturn, response)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func TestUpdateCategoryService(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedReturn model.CreateCategoryRequest
		expectedError  error
	}{{
		name: "Success Update Category",
		request: &model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:       41,
					Name:     "Mati/error",
					Priority: "Critical",
				},
				{
					Id:       42,
					Name:     "Uang tersangkut di dalam Bill acceptor",
					Priority: "High",
				},
			},
			UpdateAt: date,
		},
		expectedReturn: model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:         41,
					Name:       "Mati/error",
					IdCategory: 78,
					Priority:   "Critical",
					UpdatedAt:  date,
				},
				{
					Id:         42,
					Name:       "Uang tersangkut di dalam Bill acceptor",
					IdCategory: 78,
					Priority:   "High",
					UpdatedAt:  date,
				},
			},
			IsActive: "true",
			UpdateAt: date,
		},
		expectedError: nil,
	}}

	for _, test := range tests {
		var request_sc []*entity.SubCategory
		for _, value := range test.expectedReturn.SubCategory {
			request_sc = append(request_sc, &entity.SubCategory{
				Name:       value.Name,
				IdCategory: test.expectedReturn.Id,
				Priority:   value.Priority,
				CreatedAt:  date,
				UpdatedAt:  date,
			})
		}

		subCategoryRepository.Mock.On("DeleteSubCategory", &test.request.Id).Return(nil)
		subCategoryRepository.Mock.On("CreateSubCategory", request_sc).Return(test.expectedReturn.SubCategory, test.expectedError)
		categoryRepository.Mock.On("UpdateCategory", test.request).Return(test.expectedReturn, nil)

		t.Run(test.name, func(t *testing.T) {
			response, error := categoryService.UpdateCategory(test.request)
			require.Equal(t, test.expectedReturn, response)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func TestDeleteCategoryService(t *testing.T) {

	tests := []struct {
		name          string
		request       int
		expectedError error
	}{{
		name:          "Success Delete Category",
		request:       70,
		expectedError: nil,
	}}

	for _, test := range tests {

		categoryRepository.Mock.On("DeleteCategory", &test.request).Return(nil)

		t.Run(test.name, func(t *testing.T) {
			error := categoryService.DeleteCategory(&test.request)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func TestGetDetailCategoryService(t *testing.T) {

	date, _ := time.Parse("0001-01-01T00:00:00Z", time.RFC1123Z)
	tests := []struct {
		name              string
		request           string
		expectedReturn    []model.CreateCategoryRequest
		expectedTotalPage float64
		expectedError     error
	}{{
		name:    "Success Get Detail Category",
		request: "78",
		expectedReturn: []model.CreateCategoryRequest{
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
		categoryRepository.Mock.On("GetDetailCategory", &test.request).Return([]entity.Category{
			{
				Id:          78,
				Name:        "Bill Acceptor Rupiah",
				SubCategory: `[{"id":40,"name":"Tidak dapat menerima uang/uang tereject","idCategory":78,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":41,"name":"Mati/error","idCategory":78,"priority":"Critical","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":42,"name":"Uang tersangkut di dalam Bill acceptor","idCategory":78,"priority":"High","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]`,
				IsActive:    "true",
				UpdateAt:    date,
			},
		}, nil)

		t.Run(test.name, func(t *testing.T) {
			response, error := categoryService.GetDetailCategory(&test.request)
			require.Equal(t, test.expectedReturn, response)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func BenchmarkGetCategoryService(b *testing.B) {
	date, _ := time.Parse("0001-01-01T00:00:00Z", time.RFC1123Z)
	benchmarks := []struct {
		name    string
		request *model.GetCategoryRequest
	}{{
		name: "Benchmark Get Category",
		request: &model.GetCategoryRequest{
			Size:     100,
			PageNo:   0,
			IsActive: "true",
		},
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {
			categoryRepository.Mock.On("CountCategory", benchmark.request).Return(int(1), nil)
			categoryRepository.Mock.On("GetCategory", benchmark.request).Return([]entity.Category{
				{
					Id:          0,
					Name:        "",
					SubCategory: "",
					IsActive:    "",
					UpdateAt:    date,
				},
			}, nil)

			b.Run(benchmark.name, func(b *testing.B) {
				categoryService.GetCategory(benchmark.request)
			})
		}
	}
}

func BenchmarkCreateCategoryService(b *testing.B) {
	date := time.Now()
	benchmarks := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedReturn model.CreateCategoryRequest
	}{{
		name: "Benchmark Create Category",
		request: &model.CreateCategoryRequest{
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Name:     "Mati/error",
					Priority: "Critical",
				},
				{
					Name:     "Uang tersangkut di dalam Bill acceptor",
					Priority: "High",
				},
			},
			UpdateAt: date,
		},
		expectedReturn: model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:         41,
					Name:       "Mati/error",
					IdCategory: 78,
					Priority:   "Critical",
					UpdatedAt:  date,
				},
				{
					Id:         42,
					Name:       "Uang tersangkut di dalam Bill acceptor",
					IdCategory: 78,
					Priority:   "High",
					UpdatedAt:  date,
				},
			},
			IsActive: "true",
			UpdateAt: date,
		},
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {
			var request_sc []*entity.SubCategory
			for _, value := range benchmark.expectedReturn.SubCategory {
				request_sc = append(request_sc, &entity.SubCategory{
					Name:       value.Name,
					IdCategory: benchmark.expectedReturn.Id,
					Priority:   value.Priority,
					CreatedAt:  date,
					UpdatedAt:  date,
				})
			}

			subCategoryRepository.Mock.On("CreateSubCategory", request_sc).Return(benchmark.expectedReturn.SubCategory, nil)
			categoryRepository.Mock.On("CreateCategory", benchmark.request).Return(benchmark.expectedReturn, nil)

			b.Run(benchmark.name, func(b *testing.B) {
				categoryService.CreateCategory(benchmark.request)
			})
		}
	}
}

func BenchmarkUpdateCategoryService(b *testing.B) {
	date := time.Now()
	benchmarks := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedReturn model.CreateCategoryRequest
	}{{
		name: "Benchmark Create Category",
		request: &model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:       41,
					Name:     "Mati/error",
					Priority: "Critical",
				},
				{
					Id:       42,
					Name:     "Uang tersangkut di dalam Bill acceptor",
					Priority: "High",
				},
			},
			UpdateAt: date,
		},
		expectedReturn: model.CreateCategoryRequest{
			Id:   78,
			Name: "Bill Acceptor Rupiah",
			SubCategory: []entity.SubCategory{
				{
					Id:         41,
					Name:       "Mati/error",
					IdCategory: 78,
					Priority:   "Critical",
					UpdatedAt:  date,
				},
				{
					Id:         42,
					Name:       "Uang tersangkut di dalam Bill acceptor",
					IdCategory: 78,
					Priority:   "High",
					UpdatedAt:  date,
				},
			},
			IsActive: "true",
			UpdateAt: date,
		},
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {
			var request_sc []*entity.SubCategory
			for _, value := range benchmark.expectedReturn.SubCategory {
				request_sc = append(request_sc, &entity.SubCategory{
					Name:       value.Name,
					IdCategory: benchmark.expectedReturn.Id,
					Priority:   value.Priority,
					CreatedAt:  date,
					UpdatedAt:  date,
				})
			}

			subCategoryRepository.Mock.On("DeleteSubCategory", &benchmark.request.Id).Return(nil)
			subCategoryRepository.Mock.On("CreateSubCategory", request_sc).Return(benchmark.expectedReturn.SubCategory, nil)
			categoryRepository.Mock.On("UpdateCategory", benchmark.request).Return(benchmark.expectedReturn, nil)

			b.Run(benchmark.name, func(b *testing.B) {
				categoryService.UpdateCategory(benchmark.request)
			})
		}
	}
}

func BenchmarkDeleteCategoryService(b *testing.B) {

	benchmarks := []struct {
		name    string
		request int
	}{{
		name:    "Benchmark Delete Category",
		request: 70,
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {

			categoryRepository.Mock.On("DeleteCategory", &benchmark.request).Return(nil)

			b.Run(benchmark.name, func(b *testing.B) {
				categoryService.DeleteCategory(&benchmark.request)
			})
		}
	}
}
