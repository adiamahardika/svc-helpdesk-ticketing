package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type AreaRepositoryInterface interface {
	GetArea(request model.GetAreaRequest) ([]entity.MsArea, error)
}

func (repo *repository) GetArea(request model.GetAreaRequest) ([]entity.MsArea, error) {
	var area []entity.MsArea

	error := repo.db.Raw("SELECT * FROM ms_area WHERE area_code LIKE @AreaCode AND area_name LIKE @AreaName AND status LIKE @Status ORDER BY area_code ASC", model.GetAreaRequest{
		AreaCode: "%" + request.AreaCode + "%",
		AreaName: "%" + request.AreaName + "%",
		Status:   "%" + request.Status + "%",
	}).Find(&area).Error

	return area, error
}
