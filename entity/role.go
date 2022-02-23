package entity

type Role struct {
	Id             int          `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name"`
	ListPermission []Permission `json:"list_permission" gorm:"foreignKey:Id"`
}
