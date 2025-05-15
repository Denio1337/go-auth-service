package auth

import (
	"app/internal/service/utils"
	"app/internal/storage"
	"fmt"
)

// Clear tokens
func Logout(params *LogoutParams) error {
	const op = "service/logout"

	// Get refresh token from storage
	storageRefresh, err := storage.GetRefreshTokenByIdentity(params.Identity)
	if err != nil {
		return fmt.Errorf("%s: can not get refresh token by user identity", op)
	}

	// Compare refresh token with storage value
	isValid := utils.CompareWithHash(params.Refresh, storageRefresh.Hash)
	if !isValid {
		return fmt.Errorf("%s: identity comparison failed", op)
	}

	// Revoke refresh token
	_, err = storage.RevokeRefreshTokenByID(storageRefresh.ID)
	if err != nil {
		return fmt.Errorf("%s: can not revoke refresh token", op)
	}

	return nil
}
