package controller

import (
	"blog-api/dto"
	"blog-api/helper"
	"blog-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BlogController is a declaration contract
type BlogController interface {
	CreateBlog(ctx *gin.Context)
}

type blogController struct {
	blogService service.BlogService
}

//NewBlogController create a new instances of BlogController
func NewBlogController(blogService service.BlogService) *blogController {
	return &blogController{blogService}
}

//implement method BlogController, as a handler for the create blog method
func (c *blogController) CreateBlog(ctx *gin.Context) {
	var input dto.CreateBlog

	//Validation input user
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create blog", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//proceed to the CreateBlog method in the package service, which returns the data and error values
	newBlog, err := c.blogService.CreateBlog(input)
	if err != nil {
		response := helper.APIResponse("Failed to create blog", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create blog", http.StatusOK, "success", newBlog)
	ctx.JSON(http.StatusOK, response)
}
