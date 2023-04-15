package repository

import (
	"errors"

	"github.com/hy-reza/mygram-api/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll() ([]model.Comment, error)
	FindById(id string) (model.Comment, error)
	Save(comment model.Comment) (model.Comment, error)
	Update(comment model.Comment) (model.Comment, error)
	Delete(id string) error
}

type CommentRepositoryImpl struct {
	Db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{Db: db}
}

func (c *CommentRepositoryImpl) FindAll() ([]model.Comment, error) {
	var comments []model.Comment
	result := c.Db.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentRepositoryImpl) FindById(id string) (model.Comment, error) {
	var comment model.Comment
	result := c.Db.Find(&comment, "ID = ?", id)
	if result.RowsAffected == 0 {
		return model.Comment{}, errors.New("comment not found")
	}
	return comment, nil
}

func (c *CommentRepositoryImpl) Save(comment model.Comment) (model.Comment, error) {
	result := c.Db.Create(&comment)
	if result.Error != nil {
		return model.Comment{}, result.Error
	}
	return comment, nil
}

func (c *CommentRepositoryImpl) Update(comment model.Comment) (model.Comment, error) {
	result := c.Db.Save(&comment)
	if result.Error != nil {
		return model.Comment{}, result.Error
	}
	return comment, nil
}

func (c *CommentRepositoryImpl) Delete(id string) error {
	var comment model.Comment
	result := c.Db.Delete(&comment, "ID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
