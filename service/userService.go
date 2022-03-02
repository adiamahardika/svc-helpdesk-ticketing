package service

import (
	"encoding/json"
	"math"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]model.GetUserResponse, float64, error)
	GetUserDetail(request string) (model.GetUserResponse, error)
	DeleteUser(id int) error
}

type userService struct {
	repository repository.UserRepositoryInterface
}

func UserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (userService *userService) GetUser(request model.GetUserRequest) ([]model.GetUserResponse, float64, error) {

	var response []model.GetUserResponse
	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size
	total_data, error := userService.repository.CountUser(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.Size))

	user, error := userService.repository.GetUser(request)

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

	user, error := userService.repository.GetUserDetail(request)

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

	error := userService.repository.DeleteUser(id)

	return error
}
