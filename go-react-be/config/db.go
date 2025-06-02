package config

import (
	"log"

	"gorm.io/driver/sqlite" // Or postgres, mysql, etc.
	"gorm.io/gorm"
	"gorm.io/gorm/logger" // Import for logging queries
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	// For SQLite, just provide a file name.
	// For other databases, use connection strings.
	// Example for PostgreSQL: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(sqlite.Open("user_management.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established!")
}
