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

func (s *blogService) CreateBlog(input dto.CreateBlog) (entity.Blog, error) {
	blog := entity.Blog{}

	blog.Author = input.Author
	blog.Title = input.Title
	blog.Description = input.Description

	newBlog, err := s.repository.Save(blog)
	if err != nil {
		return newBlog, err
	}

	return newBlog, nil

}
