package repository

import "svc-myg-ticketing/entity"

type LogRepositoryInterface interface {
	CreateLog(request *entity.LgServiceActivities) error
}

func (repo *repository) CreateLog(request *entity.LgServiceActivities) error {
	error := repo.db.Table("ticketing_lg_service_activities").Create(&request).Error

	return error
}
