package controller

import (
	"blog-api/dto"
	"blog-api/helper"
	"blog-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//BlogController is a declaration contract
type BlogController interface {
	CreateBlog(ctx *gin.Context)
	FindAllBlog(ctx *gin.Context)
	FindByIDBlog(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
	DeleteByIDBlog(ctx *gin.Context)
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

//implement method BlogController, as a handler for the List all blog method
func (c *blogController) FindAllBlog(ctx *gin.Context) {
	//proceed to the FindAllBlog method in the package service, which returns the data and error values
	blogs, err := c.blogService.FindAllBlog()
	if err != nil {
		response := helper.APIResponse("Error to get blogs", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of blogs", http.StatusOK, "success", blogs)
	ctx.JSON(http.StatusOK, response)
}

//implement method BlogController, as a handler for the Detail blog method
func (c *blogController) FindByIDBlog(ctx *gin.Context) {
	//check params id if int is not
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponse("Param id not found / did not match", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//proceed to the FindByIDBlog method in the package service, which returns the data and error values
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
	} else {
		response := helper.APIResponse("List of Detail blog", http.StatusOK, "success", blog)
		ctx.JSON(http.StatusOK, response)
	}
}

//implement method BlogController, as a handler for the update blog method
func (c *blogController) UpdateBlog(ctx *gin.Context) {
	//check params id if int is not
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponse("Param id not found / did not match", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//check data blog, proceed to the FindByIDBlog method in the package service
	blog, _ := c.blogService.FindByIDBlog(id)
	if (blog.ID == 0) {
		response := helper.APIResponse("Blog not found", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	var input dto.CreateBlog
	if input.Author == "" {
		input.Author = blog.Author
	}
	if input.Title == "" {
		input.Title = blog.Title
	}
	if input.Description == "" {
		input.Description = blog.Description
	}

	//Validation input user
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update blog", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//proceed to the UpdateBlog method in the package service, which returns the data and error values
	updatedBlog, err := c.blogService.UpdateBlog(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to updated blog", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated blog", http.StatusOK, "success", updatedBlog)
	ctx.JSON(http.StatusOK, response)
}

//implement method BlogController, as a handler for the delete blog method
func (c *blogController) DeleteByIDBlog(ctx *gin.Context) {
	//check params id if int is not
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponse("Param id not found / did not match", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//check data blog, proceed to the FindByIDBlog method in the package service
	blog, _ := c.blogService.FindByIDBlog(id)
	if (blog.ID == 0) {
		response := helper.APIResponse("Blog not found", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	//proceed to the DeleteByIDBlog method in the package service, which returns the data and error values
	blog, err = c.blogService.DeleteByIDBlog(id)
	if err != nil {
		response := helper.APIResponse("Error to delete blog", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Deleted blog", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, response)
}
