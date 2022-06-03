package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type TerminalServiceInterface interface {
	GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error)
}

type terminalService struct {
	repository repository.TerminalRepositoryInterface
}

func TerminalService(repository repository.TerminalRepositoryInterface) *terminalService {
	return &terminalService{repository}
}

func (terminalService *terminalService) GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error) {

	terminal, error := terminalService.repository.GetTerminal(request)

	return terminal, error
}
