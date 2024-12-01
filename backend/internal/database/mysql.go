package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	// Construct the MySQL Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Username
		os.Getenv("DB_PASSWORD"), // Password
		os.Getenv("DB_HOST"),     // Host
		os.Getenv("DB_NAME"),     // Database name
	)

	// Open a connection to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
