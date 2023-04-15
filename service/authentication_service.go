package service

import (
	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest) (response.UserResponse, error)
}
