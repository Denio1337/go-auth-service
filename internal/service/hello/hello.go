package hello

import "app/internal/storage"

func Hello() (string, error) {
	return storage.Hello()
}
