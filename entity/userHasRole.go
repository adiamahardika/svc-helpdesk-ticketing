package entity

type UserHasRole struct {
	Id     int `json:"id" gorm:"primaryKey"`
	IdRole int `json:"idRole"`
	IdUser int `json:"idUser"`
}
