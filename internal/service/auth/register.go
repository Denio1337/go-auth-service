package auth

import (
	"app/internal/service/utils"
	"app/internal/storage"
	"app/internal/storage/model"
	"errors"
	"fmt"
)

// Add new user
func Register(params *RegisterParams) (*RegisterResult, error) {
	const op = "service/register"

	// Hash password
	hashedPassword, err := utils.Hash(params.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: can not hash password", op)
	}

	user := &model.User{
		Username: params.Username,
		Password: hashedPassword,
	}

	// Save user in storage
	err = storage.AddUser(user)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedKey) {
			return nil, ErrUserExists
		}

		return nil, fmt.Errorf("%s: can not save user in storage", op)
	}

	return &RegisterResult{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
