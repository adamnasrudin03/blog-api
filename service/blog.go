package service

import (
	"blog-api/dto"
	"blog-api/entity"
	"blog-api/repository"
)

//BlogService is a declaration contract
type BlogService interface {
	CreateBlog(input dto.CreateBlog) (entity.Blog, error)
}

type blogService struct {
	repository repository.BlogRepository
}

//NewBlogService creates an instance BlogService
func NewBlogService(repository repository.BlogRepository) *blogService {
	return &blogService{repository}
}

//implement method BlogService, func to take input from the user then save it / continue to the repository (save)
func (s *blogService) CreateBlog(input dto.CreateBlog) (entity.Blog, error) {
	blog := entity.Blog{}

	blog.Author = input.Author
	blog.Title = input.Title
	blog.Description = input.Description

	//proceed to the save method in the package repository, which returns the data and error values
	newBlog, err := s.repository.Save(blog)
	if err != nil {
		return newBlog, err
	}

	return newBlog, nil

}
