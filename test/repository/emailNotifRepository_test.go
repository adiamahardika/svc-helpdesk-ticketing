package repository_test

import (
	"regexp"
	"svc-myg-ticketing/entity"
	dbmock "svc-myg-ticketing/mock/db_mock"
	"svc-myg-ticketing/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_Repo_EmailNotif_Get(t *testing.T) {

	gormDB, mock := dbmock.DbMock(t)

	repository := repository.Repository(gormDB)

	tests := []struct {
		name           string
		expectedQuery  string
		expectedReturn []entity.EmailNotif
	}{{
		name:          "Success",
		expectedQuery: "SELECT * FROM ticketing_email_notif ORDER BY email ASC",
		expectedReturn: []entity.EmailNotif{{
			Id:        1,
			Email:     "test@mail.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		},
	}}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			sub_category := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at"}).AddRow(test.expectedReturn[0].Id, test.expectedReturn[0].Email, test.expectedReturn[0].CreatedAt, test.expectedReturn[0].UpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta(
				test.expectedQuery)).
				WillReturnRows(sub_category)

			res, err := repository.GetEmailNotif()

			require.NoError(t, err)
			require.Equal(t, test.expectedReturn, res)
		})
	}
}
