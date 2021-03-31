package main

import (
	"blog-api/config"
	"blog-api/controller"
	"blog-api/helper"
	"blog-api/repository"
	"blog-api/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDbConnection()

	blogRepository    repository.BlogRepository    = repository.NewBlogRepository(db)
	commentRepository repository.CommentRepository = repository.NewCommentRepository(db)

	blogService    service.BlogService    = service.NewBlogService(blogRepository)
	commentService service.CommentService = service.NewCommentService(commentRepository)

	blogController    controller.BlogController    = controller.NewBlogController(blogService)
	commentController controller.CommentController = controller.NewCommentController(commentService, blogService)
)

func main() {
	defer config.CloseDbConnection(db)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		response := helper.APIResponse("Welcome my application", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
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
	api.DELETE("/blogs/:id/comment/:idComment/delete", commentController.DeleteByIDComment)

	router.NoRoute(func(c *gin.Context) {
		response := helper.APIResponse("Page not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
	})

	router.Run()
}
