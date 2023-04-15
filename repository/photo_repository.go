package repository

import (
	"github.com/hy-reza/mygram-api/model"
)

type PhotosRepository interface {
	Save(photo model.Photo) (*model.Photo, error)
	FindAll() (*[]model.Photo, error)
	FindById(photoId string) (*model.Photo, error)
	Update(photo *model.Photo) (*model.Photo, error)
	Delete(photoId string) error
}
