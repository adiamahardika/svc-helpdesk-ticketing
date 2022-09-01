package model

import (
	"github.com/dgrijalva/jwt-go"
)

type LoginResponse struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	Role        []*GetRoleResponse `json:"role" gorm:"foreignKey:Id"`
	AreaCode    []string           `json:"areaCode"`
	Regional    []string           `json:"regional"`
	GrapariId   []string           `json:"grapariId"`
	AccessToken string             `json:"accessToken"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
