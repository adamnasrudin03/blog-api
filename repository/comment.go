package repository

import (
	"blog-api/entity"

	"gorm.io/gorm"
)

//CommentRepository is a declaration contract
type CommentRepository interface {
	Save(comment entity.Comment) (entity.Comment, error)
}


type commentRepository struct {
	db *gorm.DB
}

//NewCommentRepository creates an instance commentRepository
func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

//implement method CommentRepository, func to save directly to db
func (r *commentRepository) Save(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}