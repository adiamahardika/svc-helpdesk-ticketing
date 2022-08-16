package repositorymock_test

import "svc-myg-ticketing/entity"

func (repo *RepositoryMock) CreateEmailNotif(request *entity.EmailNotif) error {

	return nil
}

func (repo *RepositoryMock) GetEmailNotif() ([]entity.EmailNotif, error) {

	arguments := repo.Mock.Called()

	email_notif := arguments.Get(0).([]entity.EmailNotif)

	return email_notif, nil
}

func (repo *RepositoryMock) UpdateEmailNotif(request *entity.EmailNotif) (entity.EmailNotif, error) {

	arguments := repo.Mock.Called(request)

	email_notif := arguments.Get(0).(entity.EmailNotif)

	return email_notif, nil
}

func (repo *RepositoryMock) DeleteEmailNotif(id *int) error {

	arguments := repo.Mock.Called(id)

	email_notif := arguments.Get(0).(error)

	return email_notif
}

func (repo *RepositoryMock) GetDetailEmailNotif(id *int) ([]entity.EmailNotif, error) {

	arguments := repo.Mock.Called(id)

	email_notif := arguments.Get(0).([]entity.EmailNotif)

	return email_notif, nil
}

func (repo *RepositoryMock) GetAllEmailNotif() ([]string, error) {

	arguments := repo.Mock.Called()

	email_notif := arguments.Get(0).([]string)

	return email_notif, nil
}
