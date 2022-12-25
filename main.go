package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	"project/config"
	"project/handler"
	m "project/middleware"
	"project/repo"
	"project/routes"
	"project/services"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Print("failed to connect env")
	}
}
func main() {
	port := os.Getenv("PORT")
	//creating an instance of chi
	router := chi.NewRouter()

	config.Init()
	var (
		db              *sql.DB               = config.ConnectDB()
		userRepo        repo.UserRepository   = repo.NewUserRepo(db)
		adminRepo       repo.AdminRepository  = repo.NewAdminRepo(db)
		jwtAdminService services.JWTService   = services.NewJWTAdminService()
		jwtUserService  services.JWTService   = services.NewJWTUserService()
		adminService    services.AdminService = services.NewAdminService(adminRepo, userRepo)
		authService     services.AuthService  = services.NewAuthService(userRepo, adminRepo)
		userService     services.UserService  = services.NewUserService(userRepo)
		adminMiddleware m.Middleware          = m.NewMiddlewareAdmin(jwtAdminService)
		userMiddleware  m.Middleware          = m.NewMiddlewareUser(jwtUserService)
		authHandler     handler.AuthHandler   = handler.NewAuthHandler(jwtAdminService, jwtUserService, adminService, userService, authService)
		adminHandler    handler.AdminHandler  = handler.NewAdminHandler(adminService, userService)
		userHandler     handler.UserHandler   = handler.NewUserHandler(userService)
		userRoute       routes.UserRoute      = routes.NewUserRoute()
		adminRoute      routes.AdminRoute     = routes.NewAdminRoute()
	)

	//routing

	userRoute.UserRouter(router, userHandler, authHandler, userMiddleware)
	adminRoute.AdminRouter(router, authHandler, adminMiddleware, adminHandler)

	log.Println("Api is listening on port:", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Print("cant connect")
	}

}
