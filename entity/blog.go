package entity

import "time"

//Blog represents blogs table in database
type Blog struct {
	ID       		uint64  	`gorm:"primary_key:auto_increment" json:"id"`
	Author     		string 		`gorm:"type:varchar(255)" json:"author"`
	Title     		string 		`gorm:"type:varchar(255)" json:"title"`
	Description     string 		`gorm:"type:text" json:"description"`
	CreatedAt   	time.Time 	`json:"-"`
	UpdatedAt   	time.Time	`json:"-"`
}