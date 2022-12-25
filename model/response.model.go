package model

type AdminResponse struct {
	ID       int    `json:"id"`
	Username string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     int    `json:"role"`
	Token    string `json:"token,omitempty"`
}
type UserResponse struct {
	ID         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Token      string `json:"token"`
}
type ProductResponse struct {
	Product_name string  `json:"product_name" gorm:"not null"`
	Description  string  `json:"description" gorm:"not null"`
	Quantity     int32   `json:"quantity" gorm:"not null"`
	Image_Path   string  `json:"image_path" gorm:"not null"`
	Price        float32 `json:"price" gorm:"not null"`
	Color        string  `json:"color"`
	Available    bool    `json:"available" gorm:"not null"`
	Rating       uint    `json:"rating" gorm:"not null"`
	Trending     bool    `json:"trending" gorm:"not null"`
}
