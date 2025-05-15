package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load .env file
func init() {
	// Try to load main file
	err := godotenv.Load(".env")
	if err == nil {
		return
	}

	// Try to load example file with default values
	err = godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("error while loading .env")
	}
}

// Get .env value by key
func Get(key EnvKey) string {
	// Return default value "" if env key is invalid
	if !key.IsValid() {
		return ""
	}

	return os.Getenv(string(key))
}
