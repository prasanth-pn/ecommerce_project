package handler

import (
	"net/http"
	"project/common/response"
	"project/services"
	"project/utils"
)

type UserHandler interface {
	ViewProducts() http.HandlerFunc
}
type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{

		userService: userService,
	}
}
func (c *userHandler) ViewProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := c.userService.ViewProducts()
		if err != nil {
			respons := response.ErrorResponse("errror while getting product from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			utils.ResponseJSON(w, respons)
			return
		}
		respons := response.SuccessResponse(true, "listed all products", products)
		utils.ResponseJSON(w, respons)
	}
}
