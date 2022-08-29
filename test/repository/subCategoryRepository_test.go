package repository_test

import (
	"regexp"
	"time"

	"svc-myg-ticketing/entity"
	dbmock "svc-myg-ticketing/mock/db_mock"
	"svc-myg-ticketing/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_Repo_SubCategory_Get(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		expectedQuery  string
		expectedReturn []entity.SubCategory
	}{{
		name:          "Success",
		expectedQuery: "SELECT * FROM ticketing_sub_category ORDER BY name ASC",
		expectedReturn: []entity.SubCategory{{
			Id:         1,
			Name:       "Test",
			IdCategory: 1,
			Priority:   "High",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			sub_category := sqlmock.NewRows([]string{"id", "name", "id_category", "priority", "created_at", "updated_at"}).AddRow(test.expectedReturn[0].Id, test.expectedReturn[0].Name, test.expectedReturn[0].IdCategory, test.expectedReturn[0].Priority, test.expectedReturn[0].CreatedAt, test.expectedReturn[0].UpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta(
				test.expectedQuery)).
				WillReturnRows(sub_category)

			res, err := repository.GetSubCategory()

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}

}
