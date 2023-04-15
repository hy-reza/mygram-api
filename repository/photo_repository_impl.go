package repository

import (
	"errors"

	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"

	"gorm.io/gorm"
)

type PhotosRepositoryImpl struct {
	Db *gorm.DB
}

func NewPhotosRepositoryImpl(Db *gorm.DB) PhotosRepository {
	return &PhotosRepositoryImpl{Db: Db}
}

func (p *PhotosRepositoryImpl) Save(photo model.Photo) (*model.Photo, error) {
	result := p.Db.Create(&photo)
	return &photo, result.Error
}

func (p *PhotosRepositoryImpl) FindAll() (*[]model.Photo, error) {
	var photos []model.Photo
	results := p.Db.Preload("Comments").Find(&photos)
	helper.ErrorPanic(results.Error)

	return &photos, results.Error
}

func (p *PhotosRepositoryImpl) FindById(photoId string) (*model.Photo, error) {
	var photo model.Photo
	result := p.Db.Preload("Comments").Find(&photo, "ID = ?", photoId)
	if result.RowsAffected == 0 {
		return &model.Photo{}, errors.New("photo not found")
	}
	return &photo, nil
}

func (p *PhotosRepositoryImpl) Update(photo *model.Photo) (*model.Photo, error) {
	result := p.Db.Save(&photo)
	return photo, result.Error
}

func (p *PhotosRepositoryImpl) Delete(photoId string) error {
	result := p.Db.Delete(&model.Photo{}, "ID = ?", photoId)
	return result.Error
}
