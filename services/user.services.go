package services

import (
	"crypto/md5"
	_ "crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"project/config"
	"project/model"
	"project/repo"
	"time"
)

type UserService interface {
	FindUser(email string) (*model.UserResponse, error)
	CreateUser(newUser model.User) error
	ViewProducts() (*[]model.ProductResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
}
type userService struct {
	userRepo   repo.UserRepository
	mailConfig config.MailConfig
}

func NewUserService(
	userRepo repo.UserRepository,
	mailConfig config.MailConfig) UserService {
	return &userService{
		userRepo:   userRepo,
		mailConfig: mailConfig,
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
func (c *userService) SendVerificationEmail(email string) error {
	fmt.Println("userService")
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)
	message := fmt.Sprintf("\n the verification code is :\n\n%d.\n"+
		" Use to verify your account \n  "+
		"Thank you for using my site\n with regards team site",
		code,
	)
	fmt.Println(email)
	//send random code to user's email

	if err := c.mailConfig.SendMail(email, message); err != nil {
		return err
	}
	err := c.userRepo.StoreVerificationDetails(email, code)
	if err != nil {
		return err
	}
	return nil

}

func (c *userService) ViewProducts() (*[]model.ProductResponse, error) {

	product, err := c.userRepo.ViewProducts()

	if err != nil {
		return nil, err
	}
	return &product, nil

}
func (c *userService) VerifyAccount(email string, code int) error {

	err := c.userRepo.VerifyAccount(email, code)

	if err != nil {
		return err
	}
	return nil
}
