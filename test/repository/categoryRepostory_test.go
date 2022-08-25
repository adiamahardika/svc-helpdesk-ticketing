package repository_test

import (
	"math"
	"regexp"
	"svc-myg-ticketing/entity"
	dbmock "svc-myg-ticketing/mock/db_mock"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_Repo_Category_Get(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	date := time.Time{}
	tests := []struct {
		name           string
		request        *model.GetCategoryRequest
		expectedQuery  string
		expectedReturn []entity.Category
	}{{
		name: "Get All",
		request: &model.GetCategoryRequest{
			Size:       math.MaxInt16,
			StartIndex: 0,
			IsActive:   "",
		},
		expectedQuery: "SELECT ticketing_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_sub_category.id, 'idCategory', ticketing_sub_category.id_category, 'name', ticketing_sub_category.name, 'priority', ticketing_sub_category.priority)) AS sub_category FROM ticketing_category LEFT OUTER JOIN ticketing_sub_category ON (ticketing_category.id = ticketing_sub_category.id_category) WHERE is_active LIKE $1 GROUP BY ticketing_category.id ORDER BY name ASC LIMIT $2 OFFSET $3",
		expectedReturn: []entity.Category{
			{
				Id:          3,
				Name:        "Test Category",
				SubCategory: "list_of_sub_category",
				IsActive:    "False",
				UpdateAt:    date,
			},
		},
	}, {
		name: "Active Category",
		request: &model.GetCategoryRequest{
			Size:       100,
			StartIndex: 0,
			IsActive:   "true",
		},
		expectedQuery: "SELECT ticketing_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_sub_category.id, 'idCategory', ticketing_sub_category.id_category, 'name', ticketing_sub_category.name, 'priority', ticketing_sub_category.priority)) AS sub_category FROM ticketing_category LEFT OUTER JOIN ticketing_sub_category ON (ticketing_category.id = ticketing_sub_category.id_category) WHERE is_active LIKE $1 GROUP BY ticketing_category.id ORDER BY name ASC LIMIT $2 OFFSET $3",
		expectedReturn: []entity.Category{
			{
				Id:          3,
				Name:        "Test Category",
				SubCategory: "list_of_sub_category",
				IsActive:    "True",
				UpdateAt:    date,
			},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"id", "name", "sub_category", "is_active", "updated_at"}).AddRow(test.expectedReturn[0].Id, test.expectedReturn[0].Name, test.expectedReturn[0].SubCategory, test.expectedReturn[0].IsActive, test.expectedReturn[0].UpdateAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			res, err := repository.GetCategory(test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}

func Test_Repo_Category_Count(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		request        *model.GetCategoryRequest
		expectedQuery  string
		expectedReturn int
	}{{
		name: "Count All",
		request: &model.GetCategoryRequest{
			IsActive: "",
		},
		expectedQuery:  "SELECT COUNT(*) as total_data FROM ticketing_category WHERE is_active LIKE $1",
		expectedReturn: 10,
	}, {
		name: "Count Active",
		request: &model.GetCategoryRequest{
			Size:       100,
			StartIndex: 0,
			IsActive:   "true",
		},
		expectedQuery:  "SELECT COUNT(*) as total_data FROM ticketing_category WHERE is_active LIKE $1",
		expectedReturn: 5,
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"total_data"}).AddRow(test.expectedReturn)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			res, err := repository.CountCategory(test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}

func Test_Repo_Category_Create(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedQuery  string
		expectedReturn model.CreateCategoryRequest
	}{{
		name: "Success",
		request: &model.CreateCategoryRequest{
			Name:     "Test Category",
			IsActive: "true",
			UpdateAt: time.Time{},
		},
		expectedQuery: "INSERT INTO ticketing_category(name, is_active, update_at) VALUES($1, $2, $3) RETURNING ticketing_category.*",
		expectedReturn: model.CreateCategoryRequest{
			Id:       1,
			Name:     "Test Category",
			IsActive: "true",
			UpdateAt: time.Time{},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"id", "name", "is_active", "update_at"}).AddRow(test.expectedReturn.Id, test.expectedReturn.Name, test.expectedReturn.IsActive, test.expectedReturn.UpdateAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			res, err := repository.CreateCategory(test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}

func Test_Repo_Category_Update(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		request        *model.CreateCategoryRequest
		expectedQuery  string
		expectedReturn model.CreateCategoryRequest
	}{{
		name: "Success",
		request: &model.CreateCategoryRequest{
			Name:     "Test Category",
			UpdateAt: time.Time{},
		},
		expectedQuery: "UPDATE ticketing_category SET name = $1, update_at = $2 WHERE id = $3 RETURNING ticketing_category.*",
		expectedReturn: model.CreateCategoryRequest{
			Id:       1,
			Name:     "Test Category",
			IsActive: "true",
			UpdateAt: time.Time{},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"id", "name", "is_active", "update_at"}).AddRow(test.expectedReturn.Id, test.expectedReturn.Name, test.expectedReturn.IsActive, test.expectedReturn.UpdateAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			res, err := repository.UpdateCategory(test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}

func Test_Repo_Category_Delete(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		request        int
		expectedQuery  string
		expectedReturn model.CreateCategoryRequest
	}{{
		name:          "Success",
		request:       1,
		expectedQuery: "UPDATE ticketing_category SET is_active = $1 WHERE id = $2 RETURNING ticketing_category.*",
		expectedReturn: model.CreateCategoryRequest{
			Id:       1,
			Name:     "Test Category",
			IsActive: "true",
			UpdateAt: time.Time{},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"id", "name", "is_active", "update_at"}).AddRow(test.expectedReturn.Id, test.expectedReturn.Name, test.expectedReturn.IsActive, test.expectedReturn.UpdateAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			err := repository.DeleteCategory(&test.request)

			require.NoError(t, err)
		})
	}
}

func Test_Repo_Category_GetDetail(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	date := time.Time{}
	tests := []struct {
		name           string
		request        string
		expectedQuery  string
		expectedReturn []entity.Category
	}{{
		name:          "Success",
		request:       "3",
		expectedQuery: "SELECT ticketing_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', ticketing_sub_category.id, 'idCategory', ticketing_sub_category.id_category, 'name', ticketing_sub_category.name, 'priority', ticketing_sub_category.priority)) AS sub_category FROM ticketing_category LEFT OUTER JOIN ticketing_sub_category ON (ticketing_category.id = ticketing_sub_category.id_category) WHERE ticketing_category.id = $1 GROUP BY ticketing_category.id",
		expectedReturn: []entity.Category{
			{
				Id:          3,
				Name:        "Test Category",
				SubCategory: "list_of_sub_category",
				IsActive:    "False",
				UpdateAt:    date,
			},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			category := sqlmock.NewRows([]string{"id", "name", "sub_category", "is_active", "updated_at"}).AddRow(test.expectedReturn[0].Id, test.expectedReturn[0].Name, test.expectedReturn[0].SubCategory, test.expectedReturn[0].IsActive, test.expectedReturn[0].UpdateAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(category)

			res, err := repository.GetDetailCategory(&test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}
