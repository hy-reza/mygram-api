package repository

import (
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/model"
)

type UsersRepository interface {
	Save(users model.User) (*model.User, error)
	Update(users model.User)
	Delete(usersId int)
	FindById(usersId string) (response.UserResponse, error)
	FindAll() []response.UserResponse
	FindByUsername(username string) (model.User, error)
}
