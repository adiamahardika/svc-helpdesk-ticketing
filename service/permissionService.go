package service

import (
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
)

type PermissionServiceInterface interface {
	GetPermission() ([]*entity.Permission, error)
}

type permissionService struct {
	permissionRepository repository.PermissionRepositoryInterface
}

func PermissionService(permissionRepository repository.PermissionRepositoryInterface) *permissionService {
	return &permissionService{permissionRepository}
}

func (permissionService *permissionService) GetPermission() ([]*entity.Permission, error) {

	permission, error := permissionService.permissionRepository.GetPermission()

	return permission, error
}
