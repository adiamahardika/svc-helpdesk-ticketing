package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type GrapariRepositotyInterface interface {
	GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error)
}

func (repo *repository) GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error) {
	var grapari []entity.MsGrapari

	error := repo.db.Raw("SELECT * FROM ms_grapari WHERE regional LIKE @Regional AND area LIKE @AreaCode AND grapari_id LIKE @GrapariId AND status LIKE @Status ORDER BY name ASC", model.GetGrapariRequest{
		AreaCode:  "%" + request.AreaCode + "%",
		Regional:  "%" + request.Regional + "%",
		GrapariId: "%" + request.GrapariId + "%",
		Status:    "%" + request.Status + "%",
	}).Find(&grapari).Error

	return grapari, error
}
