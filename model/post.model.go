package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	First_Name   string `json:"first_name" gorm:"not null"`
	Last_Name    string `json:"last_name" gorm:"not null"`
	Email        string `json:"email" gorm:"not null" valid:"email"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone" gorm:"not null"`
	Password     string `json:"password" gorm:"not null" valid:"length(6/12)"`
	Status       bool   `json:"status"`
	Verification bool   `json:"verification"`
}

type Address struct {
	Name         string `json:"name" gorm:"not null"`
	State        string `json:"state" gorm:"not null"`
	Landmark     string `json:"landmark" gorm:"not null"`
	Address_Line string `json:"address_Line" gorm:"not null"`
	Phone_no     uint   `json:"phone_no" gorm:"not null"`
	Pincode      uint   `json:"pincode" gorm:"not null"`
}
type Verification struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type Admin struct {
	ID       int
	Username string
	Password string
}

type Category struct {
	gorm.Model
	Category_name string `json:"category_name" gorm:"primaryKey"`
	Image         string `json:"image" gorm:"not null"`
	Description   string `json:"description" gorm:"not null"`
}
type Product struct {
	gorm.Model

	Product_name  string  `json:"product_name" gorm:"not null"`
	Description   string  `json:"description" gorm:"not null"`
	Quantity      uint16  `json:"quantity" gorm:"not null"`
	Image_Path    string  `json:"image_path" gorm:"not null"`
	Price         float32 `json:"price" gorm:"not null"`
	Color         string  `json:"color"`
	Available     bool    `json:"available" gorm:"not null"`
	Rating        uint    `json:"rating" gorm:"not null"`
	Trending      bool    `json:"trending" gorm:"not null"`
	Category_name string  `json:"category_name"`
}

type Discount struct {
	gorm.Model
	Discount_Percentage float32   `json:"discount_percentage"`
	Expire_time         time.Time `json:"expire_time"`
}
type WishList struct {
	gorm.Model
	UserID     uint
	Product_id uint
}
type Cart struct {
	Cart_id     uint `json:"cart_id" gorm:"primaryKey"  `
	UserId      uint `json:"user_id"   `
	ProductID   uint `json:"product_id"  `
	Quantity    uint `json:"quantity" `
	Total_Price uint `json:"total_price"   `
}

type Order struct {
	gorm.Model
	Orderered_At   time.Time `json:"ordered_on"`
	Price          uint32    `json:"total_price" gorm:"not null"`
	Discount       int8      `json:"discount"`
	Payment_Method Payment   `json:"payment_method" gorm:"not null"`
}
type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
