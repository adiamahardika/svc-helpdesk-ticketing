package model

type CreateRoleHasPermissionRequest struct {
	IdRole       int `json:"idRole"`
	IdPermission int `json:"idPermission"`
}
