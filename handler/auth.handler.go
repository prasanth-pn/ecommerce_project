package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"project/common/response"
	"project/model"
	"project/services"
	"project/utils"
)

type AuthHandler interface {
	UserSignup() http.HandlerFunc
	UserLogin() http.HandlerFunc
	AdminSignup() http.HandlerFunc
	AdminLogin() http.HandlerFunc
}
type authHandler struct {
	jwtAdminService services.JWTService
	adminService    services.AdminService
	authService     services.AuthService
	jwtUserService  services.JWTService
	userService     services.UserService
}

func NewAuthHandler(
	jwtAdminService services.JWTService,
	jwtUserService services.JWTService,
	adminService services.AdminService,
	userService services.UserService,
	authService services.AuthService,
) AuthHandler {
	return &authHandler{
		jwtAdminService: jwtAdminService,
		jwtUserService:  jwtUserService,
		authService:     authService,
		userService:     userService,
		adminService:    adminService,
	}
}

func (c *authHandler) UserSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser model.User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		err := c.userService.CreateUser(newUser)
		log.Println(newUser)
		if err != nil {
			respons := response.ErrorResponse("failed to create user", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, respons)
			return
		}

		user, _ := c.userService.FindUser(newUser.Email)
		user.Password = ""

	}

}
func (c *authHandler) UserLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var userLogin model.User
		err := json.NewDecoder(r.Body).Decode(&userLogin)
		if err != nil {
			return
		}

		//verify User details
		err = c.authService.VerifyUser(userLogin.Email, userLogin.Password)

		if err != nil {
			respons := response.ErrorResponse("failed to login", err.Error(), nil)
			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, respons)
			return

		}
		//fetching user details

		user, _ := c.userService.FindUser(userLogin.Email)
		token := c.jwtUserService.GenerateToken(user.ID, user.Email, "user")
		user.Password = ""
		user.Token = token
		respons := response.SuccessResponse(true, "success", user.Token)
		utils.ResponseJSON(w, respons)
	}
}
func (c *authHandler) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var adminLogin model.Admin
		err := json.NewDecoder(r.Body).Decode(&adminLogin)
		if err != nil {
			return
		}
		//verifying admin credentials

		err = c.authService.VerifyAdmin(adminLogin.Username, adminLogin.Password)
		if err != nil {
			respons := response.ErrorResponse("failed to login", err.Error(), nil)
			w.Header().Add("Content-type", "application/json ")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, respons)
			return
		}
		//getting admin values
		admin, _ := c.adminService.FindAdmin(adminLogin.Username)
		token := c.jwtAdminService.GenerateToken(admin.ID, admin.Username, "admin")
		admin.Password = ""
		admin.Token = token
		respons := response.SuccessResponse(true, "success", admin.Token)
		utils.ResponseJSON(w, respons)
	}

}

func (c *authHandler) AdminSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin model.Admin

		//fetching data
		json.NewDecoder(r.Body).Decode(&newAdmin)
		err := c.adminService.CreateAdmin(newAdmin)
		if err != nil {
			respons := response.ErrorResponse("failed to signup", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, respons)

		}
	}
}
