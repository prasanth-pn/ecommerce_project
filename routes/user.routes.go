package routes

import (
	"github.com/go-chi/chi/v5"
	h "project/handler"
	m "project/middleware"
)

type UserRoute interface {
	UserRouter(router chi.Router,
		userHandler h.UserHandler,
		authHandler h.AuthHandler,
		userMiddleware m.Middleware,

	)
}
type userRoute struct{}

func NewUserRoute() UserRoute {
	return &userRoute{}
}

func (r *userRoute) UserRouter(routes chi.Router,
	userHandler h.UserHandler,
	authHandler h.AuthHandler,
	userMiddleware m.Middleware) {

	routes.Get("/user/view/products", userHandler.ViewProducts())

	routes.Post("/user/signup", authHandler.UserSignup())
	routes.Post("/user/login", authHandler.UserLogin())
	routes.Post("/user/send/verification", userHandler.SendVerificationEmail())
	routes.Patch("/user/verify/account", userHandler.VerifyAccount())
	routes.Group(
		func(r chi.Router) {
			r.Use(userMiddleware.AuthorizeJwt)

		})

}
