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
		db *sql.DB = config.ConnectDB()
	)

	log.Println("Api is listening on port:", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Print("cant connect")
	}

}
