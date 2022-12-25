package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	// loading parameters from env file
	databaseName := os.Getenv("DB_name")
	//formatting

	//dbUrI= fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password =%s ,
	//databaseHost,userName,databaseName,password)
	dbURI := os.Getenv("DB_source")
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	//verifying connection to the database is still alive
	err = db.Ping()
	if err != nil {
		fmt.Println("error in ping")
		log.Fatal(err)
	}
	log.Println("Connected to database ", databaseName)
	return db
}
