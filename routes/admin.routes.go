package routes

import (
	"github.com/go-chi/chi/v5"
	h "project/handler"
	m "project/middleware"
)

type AdminRoute interface {
	AdminRouter(routes chi.Router,
		authHandler h.AuthHandler,
		middleware m.Middleware,
		adminHandler h.AdminHandler)
}

type adminRoute struct{}

func NewAdminRoute() AdminRoute {
	return &adminRoute{}
}

// to hdndler admin routes
func (r *adminRoute) AdminRouter(routes chi.Router,
	authHandler h.AuthHandler,
	middleware m.Middleware,
	adminHandler h.AdminHandler,
) {
	routes.Post("/admin/signup", authHandler.AdminSignup())
	routes.Post("/admin/login", authHandler.AdminLogin())

	routes.Group(
		func(r chi.Router) {
			r.Use(middleware.AuthorizeJwt)

			r.Get("/admin/view/users", adminHandler.ViewAllUser())
			r.Get("/admin/addproduct", adminHandler.AddProducts())
		},
	)
}
