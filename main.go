package main

import (
	"blog-api/config"
	"blog-api/controller"
	"blog-api/repository"
	"blog-api/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB                 		= config.SetupDbConnection()

	blogRepository		repository.BlogRepository		= repository.NewBlogRepository(db)
	commentRepository	repository.CommentRepository	= repository.NewCommentRepository(db)

	blogService			service.BlogService				= service.NewBlogService(blogRepository)
	commentService		service.CommentService			= service.NewCommentService(commentRepository)

	blogController		controller.BlogController		= controller.NewBlogController(blogService)
	commentController	controller.CommentController	= controller.NewCommentController(commentService, blogService)
)

func main() {
	defer config.CloseDbConnection(db)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Welcome my application",
		})
	})

	//Grouping router
	api := router.Group("/api/v1")

	//Endpoint blogs
	api.POST("/blogs", blogController.CreateBlog)
	api.GET("/blogs", blogController.FindAllBlog)
	api.GET("/blogs/:id", blogController.FindByIDBlog)
	api.PUT("/blogs/:id/update", blogController.UpdateBlog)
	api.DELETE("/blogs/:id/delete", blogController.DeleteByIDBlog)

	//Endpoint comment
	api.POST("/blogs/:id/comment", commentController.CreateComment)
	api.PUT("/blogs/:id/comment/:idComment/update", commentController.UpdateComment)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
			"code":    http.StatusNotFound,
			"status":  "error",
		})
	})

	router.Run()
}