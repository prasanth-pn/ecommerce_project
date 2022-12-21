package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
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
	log.Println("Api is listening on port:", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Print("cant connect")
	}
	//c.JSON(200, gin.H{
	//	"messsage ": "ok",
	//})
}
