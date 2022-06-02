package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type RegionalRepositoryInterface interface {
	GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error)
}

func (repo *repository) GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error) {
	var regional []entity.MsRegional

	error := repo.db.Raw("SELECT * FROM ms_regional WHERE regional LIKE @Regional AND area LIKE @AreaCode AND status LIKE @Status ORDER BY regional ASC", model.GetRegionalRequest{
		AreaCode: "%" + request.AreaCode + "%",
		Regional: "%" + request.Regional + "%",
		Status:   "%" + request.Status + "%",
	}).Find(&regional).Error

	return regional, error
}
