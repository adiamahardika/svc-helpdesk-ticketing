package service

import (
	"encoding/json"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]model.GetRoleResponse, error)
}

type roleService struct {
	repository repository.RoleRepositoryInterface
}

func RoleService(repository repository.RoleRepositoryInterface) *roleService {
	return &roleService{repository}
}

func (roleService *roleService) GetRole() ([]model.GetRoleResponse, error) {
	var response []model.GetRoleResponse
	role, error := roleService.repository.GetRole()

	for _, value := range role {
		var list_permission []entity.Permission
		json.Unmarshal([]byte(value.ListPermission), &list_permission)

		response = append(response, model.GetRoleResponse{Name: value.Name, Id: value.Id, ListPermission: list_permission})
	}

	return response, error
}
