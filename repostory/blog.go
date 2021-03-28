package repostory

import (
	"gorm.io/gorm"
)


type BlogRepository interface {
}


type blogRepository struct {
	db *gorm.DB
}


func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}
