package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() // Load .env file

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string
	fmt.Print(dbUri)

	connection, error := gorm.Open("postgres", dbUri)

	if error != nil {
		fmt.Print(error)
	}

	db = connection
	db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

func GetDB() *gorm.DB {
	return db
}
