package service

import (
	"encoding/json"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]model.GetRoleResponse, error)
	CreateRole(request model.CreateRoleRequest) ([]entity.Role, error)
	UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error)
	DeleteRole(Id int) error
}

type roleService struct {
	roleRepository              repository.RoleRepositoryInterface
	roleHasPermissionRepository repository.RoleHasPermissionRepositoryInterface
}

func RoleService(roleRepository repository.RoleRepositoryInterface, roleHasPermissionRepository repository.RoleHasPermissionRepositoryInterface) *roleService {
	return &roleService{roleRepository, roleHasPermissionRepository}
}

func (roleService *roleService) GetRole() ([]model.GetRoleResponse, error) {
	var response []model.GetRoleResponse
	role, error := roleService.roleRepository.GetRole()

	for _, value := range role {
		var list_permission []entity.Permission
		json.Unmarshal([]byte(value.ListPermission), &list_permission)

		response = append(response, model.GetRoleResponse{Name: value.Name, Id: value.Id, ListPermission: list_permission})
	}

	return response, error
}

func (roleService *roleService) CreateRole(request model.CreateRoleRequest) ([]entity.Role, error) {
	var rhp_request []model.CreateRoleHasPermissionRequest
	role, error := roleService.roleRepository.CreateRole(request)

	if error == nil {
		for _, value := range request.ListPermission {
			rhp_request = append(rhp_request, model.CreateRoleHasPermissionRequest{IdRole: role[0].Id, IdPermission: value.Id})
		}
		error = roleService.roleHasPermissionRepository.CreateRoleHasPermission(rhp_request)
	}

	return role, error
}

func (roleService *roleService) UpdateRole(request model.UpdateRoleRequest) ([]model.GetRoleResponse, error) {
	var rhp_request []model.CreateRoleHasPermissionRequest
	role, error := roleService.roleRepository.UpdateRole(request)

	if error == nil {
		error = roleService.roleHasPermissionRepository.DeleteRoleHasPermission(request.Id)

		if error == nil {
			for _, value := range request.ListPermission {
				rhp_request = append(rhp_request, model.CreateRoleHasPermissionRequest{IdRole: role[0].Id, IdPermission: value.Id})
			}
			error = roleService.roleHasPermissionRepository.CreateRoleHasPermission(rhp_request)
		}

		role[0].ListPermission = request.ListPermission
	}

	return role, error
}

func (roleService *roleService) DeleteRole(Id int) error {

	error := roleService.roleRepository.DeleteRole(Id)

	return error
}