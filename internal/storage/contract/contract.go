package contract

// Storage interface
type Storage interface {
	Hello() (string, error)
}
