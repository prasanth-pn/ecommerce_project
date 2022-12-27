package repo

import (
	"database/sql"
	"log"
	"project/model"
)

type AdminRepository interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (model.AdminResponse, error)
	ViewAllUsers() ([]model.UserResponse, error)
	CreateProduct(product model.Product) error
	CreateCategory(category model.Category) error
}
type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}

}
func (c *adminRepo) CreateAdmin(admin model.Admin) error {
	query := `INSERT INTO 
admins (username,password)
VALUES ($1,$2);`
	err := c.db.QueryRow(query, admin.Username,
		admin.Password,
	).Err()
	return err
}

//	func (c *adminRepo) ViewSingleUser(user_Id int) (model.UserResponse, error) {
//		var user model.UserResponse
//		query := `SELECT
//
// id,
// first_name,
// last_name,
// email,
// phone
// FROM users WHERE id =$1;`
//
//	err := c.db.QueryRow(query,
//		user_Id).Scan(
//		&user.ID,
//		&user.First_Name,
//		&user.Last_Name,
//		&user.Email,
//		&user.Phone)
//	return user, err
//
// }
func (c *adminRepo) FindAdmin(username string) (model.AdminResponse, error) {
	log.Println("username of admin", username)
	var admin model.AdminResponse
	query :=
		`SELECT
id,
username,
password
FROM admins WHERE username=$1;`
	err := c.db.QueryRow(query, username).Scan(&admin.ID,
		&admin.Username,
		&admin.Password)
	return admin, err

}
func (c *adminRepo) ViewAllUsers() ([]model.UserResponse, error) {
	var users []model.UserResponse

	query := `SELECT id,first_name,last_name,email,gender,phone FROM users;`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Gender,
			&user.Phone)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (c *adminRepo) CreateProduct(product model.Product) error {
	query := `INSERT INTO 
products (product_name,
          description,
          quantity,
          image_path,
          price,
          color,
          available,
          rating,
          trending,
          category_name)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`

	err := c.db.QueryRow(query, product.Product_name,
		product.Description,
		product.Quantity,
		product.Image_Path,
		product.Price,
		product.Color,
		product.Available,
		product.Rating,
		product.Trending,
		product.Category_name,
	).Err()
	return err

}
func (c *adminRepo) CreateCategory(category model.Category) error {
	query := `INSERT INTO categories(category_name,image,description)
VALUES($1,$2,$3);`

	err := c.db.QueryRow(query, category.Category_name,
		category.Image,
		category.Description).Err()
	return err
}
