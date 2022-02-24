package entity

type RoleHasPermission struct {
	Id           int `json:"id" gorm:"primaryKey"`
	IdRole       int `json:"idRole"`
	IdPermission int `json:"id_permission"`
}
