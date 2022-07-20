package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type AreaRepositoryInterface interface {
	GetArea(request *model.GetAreaRequest) ([]entity.MsArea, error)
}

func (repo *repository) GetArea(request *model.GetAreaRequest) ([]entity.MsArea, error) {
	var area []entity.MsArea
	var area_code string

	if len(request.AreaCode) > 0 {
		area_code = "area_code IN @AreaCode AND"
	}

	query := fmt.Sprintf("SELECT * FROM ms_area WHERE %s area_name LIKE @AreaName AND status LIKE @Status ORDER BY area_code ASC", area_code)

	error := repo.db.Raw(query, model.GetAreaRequest{
		AreaCode: request.AreaCode,
		AreaName: "%" + request.AreaName + "%",
		Status:   "%" + request.Status + "%",
	}).Find(&area).Error

	return area, error
}
