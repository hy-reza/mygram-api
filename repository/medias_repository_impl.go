package repository

import (
	"github.com/hy-reza/mygram-api/model"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Save(media model.Media) (*model.Media, error)
	FindAll() (*[]model.Media, error)
	FindById(mediaId string) (*model.Media, error)
	Update(media *model.Media) (*model.Media, error)
	Delete(mediaId string) error
}

type mediaRepositoryImpl struct {
	db *gorm.DB
}

func NewMediaRepositoryImpl(db *gorm.DB) *mediaRepositoryImpl {
	return &mediaRepositoryImpl{db: db}
}

func (r *mediaRepositoryImpl) Save(media model.Media) (*model.Media, error) {
	err := r.db.Create(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepositoryImpl) FindAll() (*[]model.Media, error) {
	var medias []model.Media
	err := r.db.Find(&medias).Error
	if err != nil {
		return nil, err
	}
	return &medias, nil
}

func (r *mediaRepositoryImpl) FindById(mediaId string) (*model.Media, error) {
	var media model.Media
	err := r.db.Where("id = ?", mediaId).First(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepositoryImpl) Update(media *model.Media) (*model.Media, error) {
	err := r.db.Save(media).Error
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (r *mediaRepositoryImpl) Delete(mediaId string) error {
	var media model.Media
	err := r.db.Where("id = ?", mediaId).Delete(&media).Error
	return err
}
