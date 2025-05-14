package auth

import (
	"app/internal/config"
	"app/internal/service/utils"
	"app/internal/storage"
	"app/internal/storage/model"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Get access and refresh tokens
func Login(params *LoginParams) (*LoginResult, error) {
	// Getting user from storage
	user, err := storage.GetUserByUsername(params.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("service/auth: unknown user")
	}

	// Check if password is incorrect
	if isCorrectPassword := utils.CompareWithHash(params.Password, user.Password); !isCorrectPassword {
		return nil, errors.New("incorrect password")
	}

	// Generate pair id
	pairId := uuid.New().String()

	// Generate access token
	accessExpires := time.Now().Add(TokenLiveTimeAccess)
	access, err := generateToken(&TokenPayload{User: user, Identity: params.Identity}, accessExpires, pairId)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshExpires := time.Now().Add(TokenLiveTimeRefresh)
	refresh, err := generateToken(&TokenPayload{User: user, Identity: params.Identity}, refreshExpires, pairId)
	if err != nil {
		return nil, err
	}

	// Revoke previous refresh token
	err = storage.RevokeRefreshTokenByIdentity(params.Identity)
	if err != nil {
		return nil, err
	}

	// Hash refresh token
	refreshHashed, err := utils.Hash(refresh)
	if err != nil {
		return nil, err
	}

	// Save new refresh token
	err = storage.AddRefreshToken(&model.RefreshToken{
		UserID:    user.ID,
		Hash:      refreshHashed,
		PairID:    pairId,
		Identity:  params.Identity,
		ExpiresAt: refreshExpires,
	})
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Tokens: &Tokens{
			Access:  access,
			Refresh: refresh,
		},
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

// Generate access/refresh token
func generateToken(payload *TokenPayload, expires time.Time, pairId string) (string, error) {
	// Set SHA512 hash method
	access := jwt.New(jwt.SigningMethodHS512)

	// Set claims
	claims := access.Claims.(jwt.MapClaims)
	claims["username"] = payload.Username
	claims["id"] = payload.ID
	claims["pair_id"] = payload.PairID
	claims["exp"] = expires.Unix()

	// Set identity claim
	identity, err := utils.Hash(payload.Identity)
	if err != nil {
		return "", errors.New("service/auth: can't hash identity")
	}
	claims["identity"] = identity

	// Signing with secret key
	t, err := access.SignedString([]byte(config.Get("SECRET")))
	if err != nil {
		return "", fmt.Errorf("service/auth: token sign error")
	}

	return t, nil
}
