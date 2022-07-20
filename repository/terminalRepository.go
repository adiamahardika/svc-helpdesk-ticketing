package repository

import (
	"fmt"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type TerminalRepositoryInterface interface {
	GetTerminal(request *model.GetTerminalRequest) ([]entity.MsTerminal, error)
}

func (repo *repository) GetTerminal(request *model.GetTerminalRequest) ([]entity.MsTerminal, error) {
	var result []entity.MsTerminal
	var area_code string
	var regional string
	var grapari_id string
	var terminal string

	if len(request.AreaCode) > 0 {
		area_code = "area IN @AreaCode AND"
	}
	if len(request.Regional) > 0 {
		regional = "regional IN @Regional AND"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "grapari_has_terminal.grapari_id IN @GrapariId AND "
	}
	if len(request.TerminalId) > 0 {
		terminal = "ms_terminal.terminal_id IN @TerminalId AND "
	}

	query := fmt.Sprintf("SELECT ms_terminal.*, grapari_has_terminal.grapari_id AS grapari_id FROM ms_terminal LEFT OUTER JOIN grapari_has_terminal ON (ms_terminal.terminal_id = grapari_has_terminal.terminal_id)WHERE %s %s %s %s status LIKE @Status", area_code, regional, grapari_id, terminal)

	error := repo.db.Raw(query, model.GetTerminalRequest{
		TerminalId: request.TerminalId,
		GrapariId:  request.GrapariId,
		AreaCode:   request.AreaCode,
		Regional:   request.Regional,
		Status:     "%" + request.Status + "%",
	}).Find(&result).Error

	return result, error
}
