package service

import (
	"blog-api/dto"
	"blog-api/entity"
	"blog-api/repository"
	"errors"
)

//CommentService is a declaration contract
type CommentService interface {
	CreateComment(input dto.CreateComment) (entity.Comment, error)
	UpdateComment(commentID uint64, input dto.CreateComment) (entity.Comment, error)
	FindByIDComment(commentID uint64) (entity.Comment, error)
	DeleteByIDComment(commentID uint64) (entity.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
}

//NewCommentService creates an instance CommentService
func NewCommentService(commentRepository repository.CommentRepository) *commentService {
	return &commentService{commentRepository}
}

//implement method CommentService, func to take input from the user then save it / continue to the repository (save)
func (s *commentService) CreateComment(input dto.CreateComment) (entity.Comment, error) {
	comment := entity.Comment{}

	comment.Author = input.Author
	comment.Comments = input.Comments
	comment.BlogID = input.BlogID

	//proceed to the save method in the package repository, which returns the data and error values
	newComment, err := s.commentRepository.Save(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil

}

//implement method CommentService, func to checks with the FindById method in the package repository
func (s *commentService) FindByIDComment(commentID uint64) (entity.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}
	
	return comment, nil
}

//implement method CommentService, func to take input from the user then update it / continue to the repository (update)
func (s *commentService) UpdateComment(commentID uint64, input dto.CreateComment) (entity.Comment, error){
	//checks with the FindById method in the package repository, which returns data and error values
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	//check if the blogID is appropriate
	if comment.BlogID != input.BlogID {
		return comment, errors.New("not the blog comment")
	}

	comment.Author = input.Author
	comment.Comments = input.Comments
	comment.BlogID = input.BlogID

	//proceed to the Update method in the package repository, which returns the data and error values
	newComment, err := s.commentRepository.Update(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

//implement method CommentService, func to delete data by ID / proceed to repository (DeleteByID)
func (s *commentService) DeleteByIDComment(commentID uint64) (entity.Comment, error) {
	comment, err := s.commentRepository.DeleteByID(commentID)
	if err != nil {
		return comment, err
	}
	
	return comment, nil
}
