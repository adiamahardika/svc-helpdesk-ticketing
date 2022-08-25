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

	expected := []entity.SubCategory{{
		Id:         1,
		Name:       "Test",
		IdCategory: 1,
		Priority:   "High",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	},
	}

	sub_category := sqlmock.NewRows([]string{"id", "name", "id_category", "priority", "created_at", "updated_at"}).AddRow(expected[0].Id, expected[0].Name, expected[0].IdCategory, expected[0].Priority, expected[0].CreatedAt, expected[0].UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM ticketing_sub_category ORDER BY name ASC")).
		WillReturnRows(sub_category)

	res, err := repository.GetSubCategory()

	require.NoError(t, err)
	require.Equal(t, expected, res)
}
