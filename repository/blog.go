package repository

import (
	"blog-api/entity"

	"gorm.io/gorm"
)


type BlogRepository interface {
	Save(book entity.Blog) (entity.Blog, error)
}


type blogRepository struct {
	db *gorm.DB
}


func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

func (r *blogRepository) Save(blog entity.Blog) (entity.Blog, error) {
	err := r.db.Create(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}
