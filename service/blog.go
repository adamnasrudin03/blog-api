package service

import (
	"blog-api/dto"
	"blog-api/entity"
	"blog-api/repository"
)


type BlogService interface {
	CreateBlog(input dto.CreateBlog) (entity.Blog, error)
}

type blogService struct {
	repository repository.BlogRepository
}

func NewBlogService(repository repository.BlogRepository) *blogService {
	return &blogService{repository}
}

