package auth

import (
	"app/internal/service/utils"
	"app/internal/storage"
	"errors"
)

// Clear tokens
func Logout(params *LogoutParams) error {
	// Get refresh token from storage
	storageRefresh, err := storage.GetRefreshTokenByIdentity(params.Identity)
	if err != nil {
		return err
	}

	// Compare refresh token with storage value
	isValid := utils.CompareWithHash(params.Refresh, storageRefresh.Hash)
	if !isValid {
		return errors.New("identity comparison failed")
	}

	// Revoke refresh token
	err = storage.RevokeRefreshTokenByID(storageRefresh.ID)
	if err != nil {
		return err
	}

	return nil
}
