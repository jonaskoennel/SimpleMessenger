package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	hostname := os.Getenv("MESSAGES_DB_HOST")
	user := os.Getenv("MESSAGES_DB_USERNAME")
	password := os.Getenv("MESSAGES_DB_PASSWORD")
	dbname := os.Getenv("MESSAGES_DB_DATABASE")
	port := os.Getenv("MESSAGES_DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", hostname, user, password, dbname, port)
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	//return DB
}
