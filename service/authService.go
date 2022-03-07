package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(request model.LoginRequest) (model.LoginResponse, error)
}

type authService struct {
	userRepository repository.UserRepositoryInterface
	roleRepository repository.RoleRepositoryInterface
}

func AuthService(userRepository repository.UserRepositoryInterface, roleRepository repository.RoleRepositoryInterface) *authService {
	return &authService{userRepository, roleRepository}
}

func (authService *authService) Login(request model.LoginRequest) (model.LoginResponse, error) {
	var user_response model.LoginResponse
	user, error := authService.userRepository.GetUserDetail(request.Username)

	if user.Username == "" {
		error = fmt.Errorf("Username Not Found!")
	} else {
		error_check_pass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if error_check_pass != nil {
			error = fmt.Errorf("Password Not Match")
		}
		if error == nil {
			var parse_role []model.GetRoleResponse
			var role []entity.Role

			role, error = authService.roleRepository.GetDetailRole(model.GetRoleRequest{IdUser: user.Id})

			for _, value := range role {
				var list_permission []entity.Permission
				json.Unmarshal([]byte(value.ListPermission), &list_permission)

				parse_role = append(parse_role, model.GetRoleResponse{Name: value.Name, Id: value.Id, ListPermission: list_permission, IsActive: value.IsActive})
			}
			expirationTime := time.Now().Add(time.Minute * 60)
			claims := &model.Claims{
				SignatureKey: general.GetMD5Hash(request.Username, strconv.Itoa(user.Id)),
				Username:     request.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			jwtKey := []byte(os.Getenv("API_SECRET"))
			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				error = err
			}
			user_response = model.LoginResponse{
				Id:          user.Id,
				Name:        user.Name,
				Username:    user.Username,
				Email:       user.Email,
				AccessToken: tokenString,
				Role:        parse_role,
			}
		}

	}

	return user_response, error
}
