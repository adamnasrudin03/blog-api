package repostory

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
