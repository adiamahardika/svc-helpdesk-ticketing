package model

import (
	"github.com/dgrijalva/jwt-go"
)

type LoginResponse struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Username    string            `json:"username"`
	Email       string            `json:"email"`
	Role        []GetRoleResponse `json:"role" gorm:"foreignKey:Id"`
	AreaId      []string          `json:"areaId"`
	AccessToken string            `json:"accessToken"`
}

type Claims struct {
	SignatureKey string `json:"signature_key"`
	Username     string `json:"username"`
	jwt.StandardClaims
}
