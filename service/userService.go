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
	ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error)
	ResetPassword(request model.ResetPassword) (model.GetUserResponse, error)
	UpdateProfile(request model.UpdateUserRequest) (entity.User, error)
	UpdateUserStatus(request model.UpdateUserStatus) (entity.User, error)
	GetUserGroupByRole() ([]model.GetUserGroupByRoleResponse, error)
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
			Id:        value.Id,
			Username:  value.Username,
			Name:      value.Name,
			Email:     value.Email,
			Roles:     role,
			Phone:     value.Phone,
			Status:    value.Status,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			RuleId:    value.RuleId,
		})
	}

	return response, total_pages, error
}

func (userService *userService) GetUserDetail(request string) (model.GetUserResponse, error) {

	var response model.GetUserResponse

	user, error := userService.userRepository.GetUserDetail(request)

	var role []entity.Role
	json.Unmarshal([]byte(user.Roles), &role)
	var areaCode []string
	json.Unmarshal([]byte(user.AreaCode), &areaCode)
	var regional []string
	json.Unmarshal([]byte(user.Regional), &regional)
	var grapariId []string
	json.Unmarshal([]byte(user.GrapariId), &grapariId)

	response = model.GetUserResponse{
		Id:         user.Id,
		Username:   user.Username,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Status:     user.Status,
		AreaCode:   areaCode,
		Regional:   regional,
		GrapariId:  grapariId,
		Roles:      role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		TerminalId: user.TerminalId,
		RuleId:     user.RuleId,
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
				error = userService.userHasRoleRepository.CreateUserHasRole(&user.Id, &id_role)
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
		error = userService.userHasRoleRepository.DeleteUserHasRole(&request.Id)

		if error == nil {
			error = userService.userHasRoleRepository.CreateUserHasRole(&user.Id, &id_role)
		}
	}

	return user, error
}

func (userService *userService) ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error) {

	var user model.GetUserResponse
	date_now := time.Now()

	users, error := userService.userRepository.CheckUsername(request.Username)

	if len(users) < 1 {
		error = fmt.Errorf("Username Not Found!")
	} else {

		error_check_pass := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(request.OldPassword))
		if error_check_pass != nil {
			error = fmt.Errorf("Wrong Old Password!")
		} else {

			new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)

			if error_hash_pass != nil {
				error = fmt.Errorf("There was an error creating new password!")
			} else {
				request.UpdatedAt = date_now
				request.NewPassword = string(new_pass)

				user, error = userService.userRepository.ChangePassword(request)
			}
		}
	}

	return user, error
}

func (userService *userService) ResetPassword(request model.ResetPassword) (model.GetUserResponse, error) {
	var user model.GetUserResponse
	date_now := time.Now()

	users, error := userService.userRepository.CheckUsername(request.Username)

	if len(users) < 1 {
		error = fmt.Errorf("Username Not Found!")
	} else {

		new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)

		if error_hash_pass != nil {
			error = fmt.Errorf("There was an error creating new password!")
		} else {
			new_request := model.ChangePassRequest{
				Username:    request.Username,
				NewPassword: string(new_pass),
				UpdatedAt:   date_now,
			}

			user, error = userService.userRepository.ChangePassword(new_request)
		}

	}

	return user, error
}

func (userService *userService) UpdateProfile(request model.UpdateUserRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	request.UpdatedAt = date_now

	user, error := userService.userRepository.UpdateUser(request)
	user.Password = ""

	return user, error
}

func (userService *userService) UpdateUserStatus(request model.UpdateUserStatus) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	request.UpdatedAt = date_now

	user, error := userService.userRepository.UpdateUserStatus(request)
	user.Password = ""

	return user, error
}

func (userService *userService) GetUserGroupByRole() ([]model.GetUserGroupByRoleResponse, error) {
	var response []model.GetUserGroupByRoleResponse
	user, error := userService.userRepository.GetUserGroupByRole()

	for _, value := range user {
		var user_options []model.UserOptions
		json.Unmarshal([]byte(value.Options), &user_options)

		response = append(response, model.GetUserGroupByRoleResponse{
			Id:      value.Id,
			Label:   value.Label,
			Options: user_options,
		})
	}

	return response, error
}
