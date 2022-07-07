package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/general"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
	RefreshToken(context *gin.Context) (*model.LoginResponse, error)
}

type authService struct {
	userRepository repository.UserRepositoryInterface
	roleRepository repository.RoleRepositoryInterface
}

func AuthService(userRepository repository.UserRepositoryInterface, roleRepository repository.RoleRepositoryInterface) *authService {
	return &authService{userRepository, roleRepository}
}

func (authService *authService) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	var user_response *model.LoginResponse
	user, error := authService.userRepository.GetUserDetail(request.Username)

	if user.Username == "" {
		error = fmt.Errorf("Username Not Found!")
	} else if user.Status != "Active" {
		error = fmt.Errorf("This account has inactive!")
	} else {
		error_check_pass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if error_check_pass != nil {
			error = fmt.Errorf("Password Not Match")
		}
		if error == nil {
			var parse_role []*model.GetRoleResponse
			var role []*entity.Role

			role, error = authService.roleRepository.GetDetailRole(&user.RuleId)

			for _, value := range role {
				var list_permission []*entity.Permission
				json.Unmarshal([]byte(value.ListPermission), &list_permission)

				parse_role = append(parse_role, &model.GetRoleResponse{
					Name:           value.Name,
					Id:             value.Id,
					ListPermission: list_permission,
					GuardName:      value.GuardName,
				})
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
			var areaCode []string
			json.Unmarshal([]byte(user.AreaCode), &areaCode)
			var regional []string
			json.Unmarshal([]byte(user.Regional), &regional)
			var grapariId []string
			json.Unmarshal([]byte(user.GrapariId), &grapariId)

			user_response = &model.LoginResponse{
				Id:          user.Id,
				Name:        user.Name,
				Username:    user.Username,
				Email:       user.Email,
				AccessToken: tokenString,
				Role:        parse_role,
				AreaCode:    areaCode,
				Regional:    regional,
				GrapariId:   grapariId,
			}
		}

	}

	return user_response, error
}

func (authService *authService) RefreshToken(context *gin.Context) (*model.LoginResponse, error) {

	token_string := context.Request.Header.Get("token")
	claims := &model.Claims{}
	jwtKey := []byte(os.Getenv("API_SECRET"))
	var user []entity.User
	var login_response *model.LoginResponse

	decode_token, error := jwt.ParseWithClaims(token_string, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	validator_error, _ := error.(*jwt.ValidationError)

	if decode_token == nil {
		error = fmt.Errorf(fmt.Sprintf("Please provide token!"))
	} else if validator_error != nil && validator_error.Errors == jwt.ValidationErrorExpired {
		error = nil
	} else if error != nil {
		error = fmt.Errorf(fmt.Sprintf("Your token is invalid!"))
	}

	if error == nil {
		user, error = authService.userRepository.CheckUsername(claims.Username)

		expirationTime := time.Now().Add(time.Minute * 60)
		generate_token := &model.Claims{
			SignatureKey: general.GetMD5Hash(claims.Username, strconv.Itoa(user[0].Id)),
			Username:     claims.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, generate_token)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			error = err
		}

		login_response = &model.LoginResponse{
			AccessToken: tokenString,
		}
	}

	return login_response, error
}

func (authService *authService) Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {

		token_string := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse
		var responseCode string

		token, error := jwt.ParseWithClaims(token_string, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
		validator_error, _ := error.(*jwt.ValidationError)

		if token == nil {
			error = fmt.Errorf(fmt.Sprintf("Please provide token!"))
			responseCode = general.ErrorStatusCode
		} else if validator_error != nil && validator_error.Errors == jwt.ValidationErrorExpired {
			error = fmt.Errorf(fmt.Sprintf("Your token is expired!"))
			responseCode = general.ExpiredToken
		} else if error != nil {
			error = fmt.Errorf(fmt.Sprintf("Your token is invalid!"))
			responseCode = general.ErrorStatusCode
		}

		if error != nil {
			description = append(description, error.Error())
			status = model.StandardResponse{
				HttpStatusCode: http.StatusUnauthorized,
				ResponseCode:   responseCode,
				Description:    description,
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": status,
			})
			context.Abort()
		}
		context.Next()

	}
}

func (authService *authService) Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {

		signature_key := context.Request.Header.Get("signature-key")
		token_string := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse

		_, error := jwt.ParseWithClaims(token_string, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		user, error := authService.userRepository.CheckUsername(claims.Username)
		generate_sk := general.GetMD5Hash(claims.Username, strconv.Itoa(user[0].Id))

		if signature_key == "" {
			error = fmt.Errorf(fmt.Sprintf("Please provide signature-key!"))
		} else if signature_key != generate_sk {
			error = fmt.Errorf(fmt.Sprintf("Your signature-key is invalid!"))
		}

		if error != nil {
			description = append(description, error.Error())
			status = model.StandardResponse{
				HttpStatusCode: http.StatusUnauthorized,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": status,
			})
			context.Abort()
		}
		context.Next()

	}
}
