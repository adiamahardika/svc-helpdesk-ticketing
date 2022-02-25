package entity

type Role struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	IsActive       string `json:"isActive"`
	ListPermission string `json:"listPermission"`
}
