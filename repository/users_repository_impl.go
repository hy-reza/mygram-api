package repository

import (
	"errors"

	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(usersId int) {
	var users model.User
	result := u.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() []response.UserResponse {
	var users []model.User
	results := u.Db.Preload("Photos").Find(&users)
	helper.ErrorPanic(results.Error)

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = response.UserResponse{Id: user.ID.String(), Username: user.Username, Email: user.Email, Age: user.Age, Photos: user.Photos, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	}

	return userResponses
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(usersId string) (response.UserResponse, error) {
	var users model.User
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return response.UserResponse{Id: users.ID.String(), Username: users.Username, Email: users.Email, Age: users.Age, CreatedAt: users.CreatedAt, UpdatedAt: users.UpdatedAt}, nil
	} else {
		return response.UserResponse{}, errors.New("users is not found")
	}
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Save(users model.User) (*model.User, error) {
	result := u.Db.Create(&users)
	return &users, result.Error
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(users model.User) {
	var updateUsers = request.UpdateUserRequest{
		// Id:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

// FindByUsername implements UsersRepository
func (u *UsersRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var users model.User
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}
