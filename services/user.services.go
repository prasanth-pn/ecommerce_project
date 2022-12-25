package services

import (
	"crypto/md5"
	_ "crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"project/model"
	"project/repo"
)

type UserService interface {
	FindUser(email string) (*model.UserResponse, error)
	CreateUser(newUser model.User) error
	ViewProducts() (*[]model.ProductResponse, error)
}
type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(
	userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// --------------------------find user
func (c *userService) FindUser(email string) (*model.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *userService) CreateUser(newUser model.User) error {
	_, err := c.userRepo.FindUser(newUser.Email)
	if err == nil {
		return errors.New("username already exists")

	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	//hash password
	newUser.Password = HashPassword(newUser.Password)
	_, err = c.userRepo.InsertUser(newUser)
	fmt.Println("fsdffffgdfgfdg")

	if err != nil {
		return err

	}
	return nil
}

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
}

func (c *userService) ViewProducts() (*[]model.ProductResponse, error) {

	product, err := c.userRepo.ViewProducts()

	if err != nil {
		return nil, err
	}
	return &product, nil

}
