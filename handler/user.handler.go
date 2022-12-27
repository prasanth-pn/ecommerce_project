package handler

import (
	"fmt"
	"net/http"
	"project/common/response"
	"project/services"
	"project/utils"
	"strconv"
)

type UserHandler interface {
	ViewProducts() http.HandlerFunc
	SendVerificationEmail() http.HandlerFunc
	VerifyAccount() http.HandlerFunc
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

func (c *userHandler) SendVerificationEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")

		_, err := c.userService.FindUser(email)
		fmt.Println(err, "find user")
		fmt.Println(email, "userHandler")

		if err == nil {
			err = c.userService.SendVerificationEmail(email)
		}
		fmt.Println(err, "send mail")
		if err != nil {
			respons := response.ErrorResponse("Error sending verificartion mail", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, respons)
			return
		}
		respons := response.SuccessResponse(true, "verification mail sent successfully", email)
		utils.ResponseJSON(w, respons)

	}

}
func (c *userHandler) VerifyAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")
		code, _ := strconv.Atoi(r.URL.Query().Get("Code"))
		err := c.userService.VerifyAccount(email, code)
		if err != nil {
			respons := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, respons)
			return
		}
		respons := response.SuccessResponse(true, "Account verified successfully", email)
		utils.ResponseJSON(w, respons)
	}
}
