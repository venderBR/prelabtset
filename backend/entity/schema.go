package entity

import(
	"gorm.io/gorm"
)

type User struct{

	gorm.Model

	Firstname string   `valid:"required~Firstname is required"`
	Lastname  string
	StudentID string   `valid:"required~StudentID is required, matches(^[BMDbmd]\\d{7}$)~StudentID is invalid"`
	Phone     string   `valid:"required~Phone is required, stringlength(10|10)~Phone is invalid"`
	Email     string
}