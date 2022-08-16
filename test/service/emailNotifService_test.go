package service_test

import (
	"svc-myg-ticketing/entity"
	repositoryMock "svc-myg-ticketing/mock/repository_mock"
	"svc-myg-ticketing/service"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var emailNotifRepository = &repositoryMock.RepositoryMock{Mock: mock.Mock{}}
var emailNotifService = service.EmailNotifService(emailNotifRepository)

func Test_Service_EmailNotif_Create(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name           string
		request        *entity.EmailNotif
		expectedReturn entity.EmailNotif
		expectedError  error
	}{{
		name: "Success",
		request: &entity.EmailNotif{
			Email:     "devt@mail.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		expectedReturn: entity.EmailNotif{
			Id:        0,
			Email:     "devt@mail.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		expectedError: nil,
	}}

	for _, test := range tests {
		emailNotifRepository.Mock.On("CreateEmailNotif", test.request).Return(test.expectedError)

		t.Run(test.name, func(t *testing.T) {
			result, error := emailNotifService.CreateEmailNotif(test.request)
			require.Equal(t, test.expectedReturn, result)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func Test_Service_EmailNotif_Get(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name           string
		expectedReturn []entity.EmailNotif
		expectedError  error
	}{{
		name: "Success",
		expectedReturn: []entity.EmailNotif{
			{
				Id:        1,
				Email:     "dev@mail.com",
				CreatedAt: date,
				UpdatedAt: date,
			},
			{
				Id:        2,
				Email:     "dev2@gmail.com",
				CreatedAt: date,
				UpdatedAt: date,
			},
		},
		expectedError: nil,
	}}

	for _, test := range tests {
		emailNotifRepository.Mock.On("GetEmailNotif").Return(test.expectedReturn, test.expectedError)

		t.Run(test.name, func(t *testing.T) {
			result, error := emailNotifService.GetEmailNotif()
			require.Equal(t, test.expectedReturn, result)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func Test_Service_EmailNotif_Update(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name           string
		request        *entity.EmailNotif
		expectedReturn entity.EmailNotif
		expectedError  error
	}{{
		name: "Success",
		request: &entity.EmailNotif{
			Email:     "devt@mail.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		expectedReturn: entity.EmailNotif{
			Id:        0,
			Email:     "devt@mail.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		expectedError: nil,
	}}

	for _, test := range tests {
		emailNotifRepository.Mock.On("UpdateEmailNotif", test.request).Return(test.expectedReturn, test.expectedError)

		t.Run(test.name, func(t *testing.T) {
			result, error := emailNotifService.UpdateEmailNotif(test.request)
			require.Equal(t, test.expectedReturn, result)
			require.Equal(t, test.expectedError, error)
		})
	}
}
