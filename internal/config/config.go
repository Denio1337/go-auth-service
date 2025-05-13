package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load .env file
func init() {
	err := godotenv.Load(".env")
	fmt.Println("read .env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// Get .env value by key
func Get(key string) string {
	return os.Getenv(key)
}
