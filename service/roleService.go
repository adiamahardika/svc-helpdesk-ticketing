package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]entity.Role, error)
}

type roleService struct {
	repository repository.RoleRepositoryInterface
}

func RoleService(repository repository.RoleRepositoryInterface) *roleService {
	return &roleService{repository}
}

func (roleService *roleService) GetRole() ([]entity.Role, error) {
	role, error := roleService.repository.GetRole()

	return role, error
}
