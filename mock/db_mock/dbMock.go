package dbmock_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	var gormDB *gorm.DB
	sqlDB, mock, err := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})
	gormDB, err = gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return gormDB, mock
}
