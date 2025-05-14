package auth

import (
	"app/internal/storage"
	"errors"
)

// Get access and refresh tokens
func Me(params *MeParams) (*MeResult, error) {
	// Getting user from storage
	user, err := storage.GetUserByID(params.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("service/auth: unknown user")
	}

	return &MeResult{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
