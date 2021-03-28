package entity

import "time"

//Comment represents Comments table in database
type Comment struct {
	ID       		uint64  	`gorm:"primary_key:auto_increment" json:"id"`
	BlogID     		uint64  	`json:"-"`
	Author     		string 		`gorm:"type:varchar(255)" json:"author"`
	Comments     	string 		`gorm:"type:varchar(255)" json:"comments"`
	CreatedAt   	time.Time 	`json:"-"`
	UpdatedAt   	time.Time	`json:"-"`
	Blog			Blog		`gorm:"foreignkey:BlogID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}