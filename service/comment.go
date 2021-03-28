package service

import (
	"blog-api/dto"
	"blog-api/entity"
	"blog-api/repository"
)

//CommentService is a declaration contract
type CommentService interface {
	CreateComment(input dto.CreateComment) (entity.Comment, error)
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
