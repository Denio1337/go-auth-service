package contract

import (
	"app/internal/storage/model"
	"errors"
)

// Storage interface
type Storage interface {
	// Users operations
	GetUserByUsername(string) (*model.User, error)
	GetUserByID(uint) (*model.User, error)
	AddUser(*model.User) error

	// Refresh tokens operations
	GetRefreshTokenByPairID(string) (*model.RefreshToken, error)
	GetRefreshTokenByIdentity(string) (*model.RefreshToken, error)
	AddRefreshToken(*model.RefreshToken) error
	RevokeRefreshTokenByPairID(string) (uint, error)
	RevokeRefreshTokenByID(uint) (uint, error)
	RevokeRefreshTokenByIdentity(string) (uint, error)
}

var (
	ErrDuplicatedKey = errors.New("duplicated key")
)
