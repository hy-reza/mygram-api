package service

import (
	"errors"

	"github.com/hy-reza/mygram-api/config"
	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"
	"github.com/hy-reza/mygram-api/repository"
	"github.com/hy-reza/mygram-api/utils"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(u request.CreateUserRequest) (response.UserResponse, error) {

	hashedPassword, err := utils.HashPassword(u.Password)
	helper.ErrorPanic(err)

	newUser := model.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
		Age:      u.Age,
	}
	savedUser, err := a.UsersRepository.Save(newUser)

	users := response.UserResponse{
		Id:        savedUser.ID.String(),
		Username:  savedUser.Username,
		Email:     savedUser.Email,
		Age:       savedUser.Age,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
	}

	return users, err

}
