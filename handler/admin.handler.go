package handler

import (
	"dearDoctor/utils"
	"encoding/json"
	"log"
	"net/http"
	"project/common/response"
	"project/model"
	"project/services"
)

type AdminHandler interface {
	ViewAllUser() http.HandlerFunc
	ViewSingleUser() http.HandlerFunc
	AddProducts() http.HandlerFunc
}
type adminHandler struct {
	adminService services.AdminService
	userService  services.UserService
}

func (c *adminHandler) ViewSingleUser() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func NewAdminHandler(adminService services.AdminService,
	userService services.UserService) AdminHandler {
	return &adminHandler{
		userService:  userService,
		adminService: adminService,
	}
}
func (c *adminHandler) ViewAllUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := c.adminService.AllUsers()

		if err != nil {
			respons := response.ErrorResponse("error while getting user from database", err.Error(), nil)
			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, respons)
			return
		}
		respons := response.SuccessResponse(true, "Listed All users", users)
		utils.ResponseJSON(w, respons)
	}
}

func (c *adminHandler) AddProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct model.Product
		_ = json.NewDecoder(r.Body).Decode(&newProduct)
		err := c.adminService.CreateProduct(newProduct)
		log.Println(newProduct)
		if err != nil {
			respons := response.ErrorResponse("failed to create product", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, respons)
			return
		}

		//user, _ := c.userService.FindUser.Email)
		//user.Password = ""

	}
}
