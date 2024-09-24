package database

import (
	"api_restfull/src/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	DB, err = gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to the database", err)
	}

	err = DB.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatal("Failed to migrate Customer", err)
	}
	log.Println("created table customers")

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate Product", err)
	}
	log.Println("created table products")


	err = DB.AutoMigrate(&models.Sale{})
	if err != nil {
		log.Fatal("Failed to migrate Sale", err)
	}
	log.Println("created table sales")

}