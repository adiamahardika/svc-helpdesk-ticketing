package repository

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
)

type TerminalRepositoryInterface interface {
	GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error)
}

func (repo *repository) GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error) {
	var terminal []entity.MsTerminal

	error := repo.db.Raw("SELECT ms_terminal.*, grapari_has_terminal.grapari_id AS grapari_id FROM ms_terminal LEFT OUTER JOIN grapari_has_terminal ON (ms_terminal.terminal_id = grapari_has_terminal.terminal_id)WHERE ms_terminal.terminal_id LIKE @TerminalId AND grapari_has_terminal.grapari_id LIKE @GrapariId AND area LIKE @AreaCode AND regional LIKE @Regional AND status LIKE @Status", model.GetTerminalRequest{
		TerminalId: "%" + request.TerminalId + "%",
		GrapariId:  "%" + request.GrapariId + "%",
		AreaCode:   "%" + request.AreaCode + "%",
		Regional:   "%" + request.Regional + "%",
		Status:     "%" + request.Status + "%",
	}).Find(&terminal).Error

	return terminal, error
}
