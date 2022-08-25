package repository

import (
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

func Test_Repo_Area_Get(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	date := time.Now()
	tests := []struct {
		name           string
		request        *model.GetAreaRequest
		expectedQuery  string
		expectedReturn []entity.MsArea
	}{{
		name: "All Area",
		request: &model.GetAreaRequest{
			AreaCode: []string{},
			AreaName: "",
			Status:   "",
		},
		expectedQuery: "SELECT * FROM ms_area WHERE area_name LIKE $1 AND status LIKE $2 ORDER BY area_code ASC",
		expectedReturn: []entity.MsArea{
			{
				Id:        3,
				AreaCode:  "1",
				AreaName:  "Sumatra",
				Status:    "A",
				CreatedAt: date,
				UpdatedAt: date,
			},
		},
	}, {
		name: "Filled Request",
		request: &model.GetAreaRequest{
			AreaCode: []string{"1"},
			AreaName: "Sumatra",
			Status:   "A",
		},
		expectedQuery: "SELECT * FROM ms_area WHERE area_code IN ($1) AND area_name LIKE $2 AND status LIKE $3 ORDER BY area_code ASC",
		expectedReturn: []entity.MsArea{
			{
				Id:        3,
				AreaCode:  "1",
				AreaName:  "Sumatra",
				Status:    "A",
				CreatedAt: date,
				UpdatedAt: date,
			},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			area := sqlmock.NewRows([]string{"id", "area_code", "area_name", "status", "created_at", "updated_at"}).AddRow(test.expectedReturn[0].Id, test.expectedReturn[0].AreaCode, test.expectedReturn[0].AreaName, test.expectedReturn[0].Status, test.expectedReturn[0].CreatedAt, test.expectedReturn[0].UpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta(test.expectedQuery)).WillReturnRows(area)

			res, err := repository.GetArea(test.request)

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}
