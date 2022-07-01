package entity

type Role struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	GuardName      string `json:"guardName"`
	ListPermission string `json:"listPermission"`
}
