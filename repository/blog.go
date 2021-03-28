package repository

import (
	"blog-api/entity"

	"gorm.io/gorm"
)

//BlogRepository is a declaration contract
type BlogRepository interface {
	Save(book entity.Blog) (entity.Blog, error)
}


type blogRepository struct {
	db *gorm.DB
}

//NewBlogRepository creates an instance BlogRepository
func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

//implement method BlogRepository, func to save directly to db
func (r *blogRepository) Save(blog entity.Blog) (entity.Blog, error) {
	err := r.db.Create(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}
