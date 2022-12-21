package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         int    `json:"user_id" gorm:"primary_key"`
	First_Name string `json:"first_name" gorm:"not null"`
	Last_Name  string `json:"last_name" gorm:"not null"`
	Email      string `json:"email" gorm:"not null" valid:"email"`
	Phone      int64  `json:"phone" gorm:"not null"`
	Password   string `json:"password" gorm:"not null" valid:"length(6/12)"`
}
type Admin struct {
	ID       int
	Username string
	Password string
}
