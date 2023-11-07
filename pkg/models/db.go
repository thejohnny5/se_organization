package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func ConnectClient() (*DBClient, error) {
	// PostgreSQL connection string
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	os.Getenv("DB_HOST"),     // e.g. "localhost"
	// 	os.Getenv("DB_USER"),     // e.g. "myuser"
	// 	os.Getenv("DB_PASSWORD"), // e.g. "mypassword"
	// 	os.Getenv("DB_NAME"),     // e.g. "mydbname"
	// 	os.Getenv("DB_PORT"),     // e.g. "5432"
	// )
	dsn := "host=localhost user=postgres dbname=test port=5432"
	// Initialize GORM DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	client := &DBClient{DB: db}
	return client, nil
}
