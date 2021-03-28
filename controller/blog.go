package controller

import "blog-api/service"

type BlogController interface {

}

type blogController struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) *blogController {
	return &blogController{blogService}
}
