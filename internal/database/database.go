package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectStr := fmt.Sprintf("host=%s port=%s user =%s dbname=%s password%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectStr)
	if err != nil {
		return db, nil
	}

	if err := db.DB().Ping; err != nil {
		return db, err
	}


	return nil, nil
}
