package auth

import (
	"app/internal/service/utils"
	"app/internal/storage"
	"app/internal/storage/model"
)

func Register(params *RegisterParams) (*RegisterResult, error) {
	hashedPassword, err := utils.Hash(params.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: params.Username,
		Password: hashedPassword,
	}

	// Adding user to storage
	err = storage.AddUser(user)
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
