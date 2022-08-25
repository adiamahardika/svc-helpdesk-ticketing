package repository_test

import (
	"database/sql"
	"regexp"
	"time"

	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repository.SubCategoryRepositoryInterface
}

func (s *Suite) SetupSuite() {
	var db *sql.DB
	var err error
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = repository.Repository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func Test_Rep_Init(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Repo_SubCategory_Get() {
	expected := []entity.SubCategory{{
		Id:         1,
		Name:       "Test",
		IdCategory: 1,
		Priority:   "High",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	},
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM ticketing_sub_category ORDER BY name ASC")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "id_category", "priority", "created_at", "updated_at"}).AddRow(expected[0].Id, expected[0].Name, expected[0].IdCategory, expected[0].Priority, expected[0].CreatedAt, expected[0].UpdatedAt))

	res, err := s.repository.GetSubCategory()

	require.NoError(s.T(), err)
	require.Equal(s.T(), expected, res)
}
