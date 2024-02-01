package entity

import(
	"gorm.io/gorm"
)

type User struct{

	gorm.Model

	Firstname string   `valid:"required~Firstname is required"`
	Lastname  string
	StudentID string   `valid:"required~StudentID is required, matches(^[BMDdmd]\\{d7}$)~StudentID is invalid"`
	Phone     string
	Email     string
}