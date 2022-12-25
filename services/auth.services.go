package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	"project/repo"
)

type AuthService interface {
	VerifyUser(email string, password string) error
	VerifyAdmin(email string, password string) error
}
type authService struct {
	userRepo  repo.UserRepository
	adminRepo repo.AdminRepository
}

func NewAuthService(userRepo repo.UserRepository, adminRepo repo.AdminRepository) AuthService {
	return &authService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}
func (c *authService) VerifyUser(email string, password string) error {
	user, err := c.userRepo.FindUser(email)
	if err != nil {
		return errors.New("failed to login. check your email")

	}
	isValidPassword := VerifyPassword(password, user.Password)

	if !isValidPassword {
		return errors.New("failed to login check your credential ")

	}
	return nil
}

func VerifyPassword(requestPassword, dbPassword string) bool {
	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword

}
func (c *authService) VerifyAdmin(email, password string) error {
	admin, err := c.adminRepo.FindAdmin(email)
	if err != nil {
		return errors.New("invalid username/password, failed to login")
	}
	isValidPassword := VerifyPassword(password, admin.Password)
	if !isValidPassword {
		return errors.New("invalid username/Password,failed to login ")

	}
	return nil
}
