package dto

//CreateBlog is used when client post
type CreateBlog struct {
	Author     		string 		`json:"author" binding:"required"`
	Title     		string 		`json:"title" binding:"required"`
	Description     string 		`json:"description" binding:"required"`
}
