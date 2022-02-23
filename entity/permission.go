package entity

type Permission struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	PermissionCode string `json:"code"`
}
