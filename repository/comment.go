package repository

import (
	"blog-api/entity"

	"gorm.io/gorm"
)

//CommentRepository is a declaration contract
type CommentRepository interface {
	Save(comment entity.Comment) (entity.Comment, error)
	FindByID(commentID uint64) (entity.Comment, error)
	Update(comment entity.Comment) (entity.Comment, error) 
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

//implement method CommentRepository, func to get by id comment directly to db
func (r *commentRepository) FindByID(commentID uint64) (entity.Comment, error){
	var comment entity.Comment

	err := r.db.Where("id = ?", commentID).Find(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

//implement method CommentRepository, func to update comment directly to db
func (r *commentRepository) Update(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Save(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}
