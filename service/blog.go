package service

import (
	"blog-api/dto"
	"blog-api/entity"

	"gorm.io/gorm"
)


type BlogService interface {
	Save(input dto.CreateBlog) (entity.Blog, error)
}

type blogService struct {
	db *gorm.DB
}

func NewBlogService(db *gorm.DB) *blogService {
	return &blogService{db}
}