package controller

import (
	"blog-api/dto"
	"blog-api/helper"
	"blog-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CommentController is a declaration contract
type CommentController interface {
	CreateComment(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	blogService service.BlogService
}

//NewCommentController create a new instances of CommentController
func NewCommentController(commentService service.CommentService, blogService service.BlogService) *commentController {
	return &commentController{commentService, blogService}
}

//implement method CommentController, as a handler for the create Comment method
func (c *commentController) CreateComment(ctx *gin.Context) {
	var input dto.CreateComment
	
	//check params id if int is not
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponse("Param id not found / did not match", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	//proceed to the FindByIDBlog method in the package blogService
	blog, err := c.blogService.FindByIDBlog(id)
	if err != nil {
		response := helper.APIResponse("Error to get blog", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	//check the data returned
	if (blog.ID == 0) {
		response := helper.APIResponse("Blog not found", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	input.BlogID = id

	//Validation input user
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create comment", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//proceed to the CreateComment method in the package service, which returns the data and error values
	newComment, err := c.commentService.CreateComment(input)
	if err != nil {
		response := helper.APIResponse("Failed to create comment", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create comment", http.StatusOK, "success", newComment)
	ctx.JSON(http.StatusOK, response)
}
