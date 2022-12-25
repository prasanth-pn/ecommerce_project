package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"project/model"
)

func Init() *gorm.DB {

	dbURL := os.Getenv("DB_source")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Admin{})
	if err != nil {
		log.Println("error is happen ")
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("error is happen user")
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		log.Println("ERROR IS HAPPEN WHILE IN FETCHING PRODUCT DATA")
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		log.Println("Error is while in fetching category data")
	}
	return db
}
