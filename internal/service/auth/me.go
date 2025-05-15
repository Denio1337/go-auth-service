package auth

import (
	"app/internal/storage"
	"fmt"
)

// Get user info
func Me(params *MeParams) (*MeResult, error) {
	const op = "service/me"

	// Get user from storage
	user, err := storage.GetUserByID(params.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: can not get user from storage", op)
	}
	// User not found
	if user == nil {
		return nil, fmt.Errorf("%s: unknown username", op)
	}

	return &MeResult{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
