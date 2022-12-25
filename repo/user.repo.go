package repo

import (
	"database/sql"
	"project/model"
)

type UserRepository interface {
	FindUser(email string) (model.UserResponse, error)
	InsertUser(user model.User) (int, error)
	ViewProducts() ([]model.ProductResponse, error)
}
type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (c *userRepo) FindUser(email string) (model.UserResponse, error) {
	var user model.UserResponse

	query := `SELECT
    id,
first_name,
last_name,
email,
gender,
password,
phone
FROM users
WHERE email=$1;`

	err := c.db.QueryRow(query, email).Scan(&user.ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Gender,
		&user.Password,
		&user.Phone,
	)
	return user, err
}

func (c *userRepo) InsertUser(user model.User) (int, error) {
	var id int
	query := `INSERT INTO users(
                  first_name,
                  last_name,
                  email,
                  gender,
                  phone,
                  password)
    VALUES($1,$2,$3,$4,$5,$6)
RETURNING id;`
	err := c.db.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Gender,
		user.Phone,
		user.Password).Scan(&id)
	return id, err

}
func (c *userRepo) ViewProducts() ([]model.ProductResponse, error) {
	var products []model.ProductResponse

	query := `SELECT 
product_name,
description,
image_path,
price,
color,
available,
rating FROM products;`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.ProductResponse
		err = rows.Scan(&product.Product_name,
			&product.Description,
			&product.Image_Path,
			&product.Price,
			&product.Color,
			&product.Available,
			&product.Rating)

		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}
