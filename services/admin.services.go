package services

import (
	"database/sql"
	"errors"
	"log"
	"project/model"
	"project/repo"
)

type AdminService interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (*model.AdminResponse, error)
	//AllUsers(pagenation utils.Filter) (*[]model.UserResponse, utils.Metadata)Metadata
	AllUsers() (*[]model.UserResponse, error)
	CreateProduct(product model.Product) error
}
type adminService struct {
	adminRepo repo.AdminRepository
	userRepo  repo.UserRepository
}

func NewAdminService(
	adminRepo repo.AdminRepository,
	userRepo repo.UserRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}
func (c *adminService) FindAdmin(username string) (*model.AdminResponse, error) {
	admin, err := c.adminRepo.FindAdmin(username)
	if err != nil {
		return nil, err

	}
	return &admin, nil
}

func (c *adminService) CreateAdmin(admin model.Admin) error {
	_, err := c.adminRepo.FindAdmin(admin.Username)
	if err == nil {
		return errors.New("Admin already exists")

	}
	if err != nil && err != sql.ErrNoRows {
		return err

	}
	//hashing the password
	admin.Password = HashPassword(admin.Password)
	err = c.adminRepo.CreateAdmin(admin)
	if err != nil {
		log.Println(err)
		return errors.New("error while signup")

	}
	return nil

}

func (c *adminService) AllUsers() (*[]model.UserResponse, error) {
	users, err := c.adminRepo.ViewAllUsers()
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (c *adminService) CreateProduct(product model.Product) error {
	err := c.adminRepo.CreateProduct(product)
	if err != nil {
		log.Println(err)
		return errors.New("error while signup")

	}
	return nil

}
