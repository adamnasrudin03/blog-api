package dto

//CreateComment is used when client post
type CreateComment struct {
	BlogID     		uint64  	`json:"blog_id" binding:"required"`
	Author     		string 		`json:"author" binding:"required"`
	Comments     	string 		`json:"comments" binding:"required"`
}
