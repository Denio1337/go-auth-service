package storage

import (
	"app/internal/storage/contract"
	"app/internal/storage/impl/postgres"
)

// Global storage instance
var instance contract.Storage

// Create DB Connection with current implementation
func MustConnect() {
	var err error
	instance, err = postgres.New() // Here we can switch implementation
	if err != nil {
		panic("Can't create connection to database")
	}
}

// Interface

func Hello() (string, error) {
	return instance.Hello()
}
