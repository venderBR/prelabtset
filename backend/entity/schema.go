package entity

import(
	"gorm.io/gorm"
)

type User struct{

	gorm.Model

	Firstname string   `valid:"required~Firstname is required"`
	Lastname  string
	StudentID string
	Phone     string
	Email     string
}