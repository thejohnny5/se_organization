package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func ConnectClient() (*DBClient, error) {
	// PostgreSQL connection string
	var dsn string
	if os.Getenv("mode") == "prod" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("POSTGRES_HOST"),     // e.g. "localhost"
			os.Getenv("POSTGRES_USER"),     // e.g. "myuser"
			os.Getenv("POSTGRES_PASSWORD"), // e.g. "mypassword"
			os.Getenv("POSTGRES_DB"),       // e.g. "mydbname"
			os.Getenv("DB_PORT"),           // e.g. "5432"
		)
	} else {
		dsn = "host=localhost user=postgres dbname=test port=5432"
	}

	// Initialize GORM DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	client := &DBClient{DB: db}
	return client, nil
}
