package service

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]model.GetUserResponse, float64, error)
	GetUserDetail(request string) (model.GetUserResponse, error)
	DeleteUser(id int) error
	CreateUser(request model.CreateUserRequest) (entity.User, error)
	UpdateUser(request model.UpdateUserRequest) (entity.User, error)
}

type userService struct {
	userRepository        repository.UserRepositoryInterface
	userHasRoleRepository repository.UserHasRoleRepositoryInterface
}

func UserService(userRepository repository.UserRepositoryInterface, userHasRoleRepository repository.UserHasRoleRepositoryInterface) *userService {
	return &userService{userRepository, userHasRoleRepository}
}

func (userService *userService) GetUser(request model.GetUserRequest) ([]model.GetUserResponse, float64, error) {

	var response []model.GetUserResponse
	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size
	total_data, error := userService.userRepository.CountUser(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.Size))

	user, error := userService.userRepository.GetUser(request)

	for _, value := range user {
		var role []entity.Role
		json.Unmarshal([]byte(value.Roles), &role)

		response = append(response, model.GetUserResponse{
			Id:         value.Id,
			Username:   value.Username,
			Name:       value.Name,
			Email:      value.Email,
			Password:   "",
			Area:       value.Area,
			Roles:      role,
			Regional:   value.Regional,
			Phone:      value.Phone,
			Status:     value.Status,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
			TerminalId: value.TerminalId,
			RuleId:     value.RuleId,
			GrapariId:  value.GrapariId,
		})
	}

	return response, total_pages, error
}

func (userService *userService) GetUserDetail(request string) (model.GetUserResponse, error) {

	var response model.GetUserResponse

	user, error := userService.userRepository.GetUserDetail(request)

	var role []entity.Role
	json.Unmarshal([]byte(user.Roles), &role)

	response = model.GetUserResponse{
		Id:         user.Id,
		Username:   user.Username,
		Name:       user.Name,
		Email:      user.Email,
		Password:   "",
		Phone:      user.Phone,
		Status:     user.Status,
		Area:       user.Area,
		Roles:      role,
		Regional:   user.Regional,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		TerminalId: user.TerminalId,
		RuleId:     user.RuleId,
		GrapariId:  user.GrapariId,
	}

	return response, error
}

func (userService *userService) DeleteUser(id int) error {

	error := userService.userRepository.DeleteUser(id)

	return error
}

func (userService *userService) CreateUser(request model.CreateUserRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	users, error := userService.userRepository.CheckUsername(request.Username)

	if len(users) > 0 {
		error = fmt.Errorf("Username already exist!")
	} else {
		new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		if error_hash_pass != nil {
			error = fmt.Errorf("There was an error creating new password!")
		} else {

			request.CreatedAt = date_now
			request.UpdatedAt = date_now
			request.Password = string(new_pass)
			request.Status = "Active"
			id_role, _ := strconv.Atoi(request.Role)

			user, error = userService.userRepository.CreateUser(request)
			user.Password = ""
			if error == nil {
				error = userService.userHasRoleRepository.CreateUserHasRole(user.Id, id_role)
			}
		}
	}

	return user, error
}

func (userService *userService) UpdateUser(request model.UpdateUserRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	request.UpdatedAt = date_now
	id_role, _ := strconv.Atoi(request.Role)

	user, error := userService.userRepository.UpdateUser(request)
	user.Password = ""

	if error == nil {
		error = userService.userHasRoleRepository.DeleteUserHasRole(request.Id)

		if error == nil {
			error = userService.userHasRoleRepository.CreateUserHasRole(user.Id, id_role)
		}
	}

	return user, error
}
