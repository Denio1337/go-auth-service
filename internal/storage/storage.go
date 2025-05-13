package storage

import (
	"app/internal/storage/contract"
	"app/internal/storage/impl"
)

// Global storage instance
var instance contract.Storage

// Create DB Connection with current implementation
func init() {
	instance = impl.Impl
}

// Interface

func Hello() (string, error) {
	return instance.Hello()
}
