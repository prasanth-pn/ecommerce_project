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
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate((&model.User{}))
	return db
}
