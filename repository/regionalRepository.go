package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type RegionalRepositoryInterface interface {
	GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error)
}

func (repo *repository) GetRegional(request model.GetRegionalRequest) ([]entity.MsRegional, error) {
	var result []entity.MsRegional
	var area_code string
	var regional string

	if len(request.AreaCode) > 0 {
		area_code = "area IN @AreaCode AND "
	}

	if len(request.Regional) > 0 {
		regional = "regional IN @Regional AND "
	}

	query := fmt.Sprintf("SELECT * FROM ms_regional WHERE %s %s status LIKE @Status ORDER BY regional ASC", area_code, regional)

	error := repo.db.Raw(query, model.GetRegionalRequest{
		AreaCode: request.AreaCode,
		Regional: request.Regional,
		Status:   "%" + request.Status + "%",
	}).Find(&result).Error

	return result, error
}
