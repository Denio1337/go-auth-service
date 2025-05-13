package postgres

import (
	"app/internal/config"
	"app/internal/storage/contract"
	"app/internal/storage/model"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage implementation structure
type PostgresStorage struct {
	db *gorm.DB
}

// Create new PostgreSQL storage implementation
func New() (contract.Storage, error) {
	// Parse port from environment
	p := config.Get("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	// Define data source name
	dsn := fmt.Sprintf(
		"host=db port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
	)

	// Trying to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection opened")

	// Migrate schemas to database
	db.AutoMigrate(&model.User{})
	fmt.Println("Database migrated")

	return &PostgresStorage{db: db}, nil
}

// Interface

func (s *PostgresStorage) Hello() (string, error) {
	return "Hello", nil
}
