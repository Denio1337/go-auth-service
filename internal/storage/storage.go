package storage

import (
	"app/internal/storage/contract"
	"app/internal/storage/impl"
	"app/internal/storage/model"
)

// Global storage instance
var instance contract.Storage

// Create DB Connection with current implementation
func init() {
	instance = impl.Impl
}

// Interface

func GetUserByUsername(username string) (*model.User, error) {
	return instance.GetUserByUsername(username)
}

func AddUser(user *model.User) error {
	return instance.AddUser(user)
}

func GetUserByID(id uint) (*model.User, error) {
	return instance.GetUserByID(id)
}

func GetRefreshTokenByPairID(hash string) (*model.RefreshToken, error) {
	return instance.GetRefreshTokenByPairID(hash)
}

func AddRefreshToken(token *model.RefreshToken) error {
	return instance.AddRefreshToken(token)
}

func GetRefreshTokenByIdentity(identity string) (*model.RefreshToken, error) {
	return instance.GetRefreshTokenByIdentity(identity)
}

func RevokeRefreshTokenByIdentity(identity string) error {
	return instance.RevokeRefreshTokenByIdentity(identity)
}

func RevokeRefreshTokenByID(id uint) error {
	return instance.RevokeRefreshTokenByID(id)
}
