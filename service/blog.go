package service

import (
	"blog-api/dto"
	"blog-api/entity"
	"blog-api/repository"
)

//BlogService is a declaration contract
type BlogService interface {
	CreateBlog(input dto.CreateBlog) (entity.Blog, error)
	FindAllBlog() ([]entity.Blog, error)
	FindByIDBlog(blogID uint64) (entity.Blog, error)
	UpdateBlog(blogID uint64, input dto.CreateBlog)  (entity.Blog, error)
	DeleteByIDBlog(blogID uint64) (entity.Blog, error)
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

//implement method BlogService, func to Search for all data / continue to the repository (findAll)
func (s *blogService) FindAllBlog() ([]entity.Blog, error) {
	Blogs, err := s.repository.FindAll()
	if err != nil {
		return Blogs, err
	}
	
	return Blogs, nil
}

//implement method BlogService, func to checks with the FindById method in the package repository
func (s *blogService) FindByIDBlog(blogID uint64) (entity.Blog, error) {
	Blog, err := s.repository.FindByID(blogID)
	if err != nil {
		return Blog, err
	}
	
	return Blog, nil
}

//implement method BlogService, func to take input from the user then update it / continue to the repository (update)
func (s *blogService) UpdateBlog(blogID uint64, input dto.CreateBlog) (entity.Blog, error){
	//checks with the FindById method in the package repository, which returns data and error values
	blog, err := s.repository.FindByID(blogID)
	if err != nil {
		return blog, err
	}

	blog.Author = input.Author
	blog.Title = input.Title
	blog.Description = input.Description

	//proceed to the Update method in the package repository, which returns the data and error values
	newBlog, err := s.repository.Update(blog)
	if err != nil {
		return newBlog, err
	}

	return newBlog, nil
}

//implement method BlogService, func to delete data by ID / proceed to repository (DeleteByID)
func (s *blogService) DeleteByIDBlog(BlogID uint64) (entity.Blog, error) {
	Blog, err := s.repository.DeleteByID(BlogID)
	if err != nil {
		return Blog, err
	}
	
	return Blog, nil
}
