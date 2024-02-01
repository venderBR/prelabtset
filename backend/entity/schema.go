package entity

import(
	"gorm.io/gorm"
)

type User struct{

	gorm.Model

	Firstname string   
	Lastname  string
	StudentID string
	Phone     string
	Email     string
}