package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"project/common/response"
	"project/services"
	"project/utils"
	"strings"
)

type Middleware interface {
	AuthorizeJwt(http.Handler) http.Handler
}

type middleware struct {
	jwtService services.JWTService
}

func NewMiddlewareAdmin(jwtAdminService services.JWTService) Middleware {
	return &middleware{
		jwtService: jwtAdminService,
	}
}
func NewMiddlewareUser(jwtUserService services.JWTService) Middleware {
	return &middleware{
		jwtService: jwtUserService,
	}

}
func (c *middleware) AuthorizeJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		bearerToken := strings.Split(authHeader, " ")
		fmt.Println(bearerToken)

		fmt.Println(len(bearerToken))
		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			respons := response.ErrorResponse("failed to authenticate jwt", err.Error(), nil)
			w.Header().Add("Context-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, respons)
			return
		}
		authtoken := bearerToken[1]
		ok, claims := c.jwtService.VerifyToken(authtoken)
		if !ok {
			err := errors.New("your token is not valid")
			respons := response.ErrorResponse("Error", err.Error(), nil)
			w.Header().Add("Context-Type", "application/json")
			utils.ResponseJSON(w, respons)
			return
		}
		user_email := fmt.Sprint(claims.UserName)
		r.Header.Set("email", user_email)
		next.ServeHTTP(w, r)

	})
}
