package auth

import (
	"app/internal/config"
	"app/internal/service/utils"
	"app/internal/storage"
	"app/internal/storage/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Get access and refresh tokens
func Login(params *LoginParams) (*LoginResult, error) {
	const op = "service/login"

	// Get user from storage
	user, err := storage.GetUserByUsername(params.Username)
	if err != nil {
		return nil, fmt.Errorf("%s: can not get user from storage", op)
	}
	// User not found
	if user == nil {
		return nil, fmt.Errorf("%s: unknown username", op)
	}

	// Check if password is incorrect
	if isCorrectPassword := utils.CompareWithHash(params.Password, user.Password); !isCorrectPassword {
		return nil, fmt.Errorf("%s: incorrect user password", op)
	}

	// Generate general tokens parameters
	pairId := uuid.New().String()
	payload := &TokenPayload{
		Username: user.Username,
		ID:       user.ID,
		Identity: params.Identity,
		PairID:   pairId,
	}

	// Generate access token
	accessExpires := time.Now().Add(TokenLiveTimeAccess)
	access, err := generateToken(payload, accessExpires)
	if err != nil {
		return nil, fmt.Errorf("%s: can not generate access token", op)
	}

	// Generate refresh token
	refreshExpires := time.Now().Add(TokenLiveTimeRefresh)
	refresh, err := generateToken(payload, refreshExpires)
	if err != nil {
		return nil, fmt.Errorf("%s: can not generate refresh token", op)
	}

	// Revoke previous refresh token
	_, err = storage.RevokeRefreshTokenByIdentity(params.Identity)
	if err != nil {
		return nil, fmt.Errorf("%s: can not revoke refresh token", op)
	}

	// Hash refresh token
	refreshHashed, err := utils.Hash(refresh)
	if err != nil {
		return nil, fmt.Errorf("%s: can not hash refresh token", op)
	}

	// Save new refresh token
	err = storage.AddRefreshToken(&model.RefreshToken{
		UserID:    user.ID,
		Hash:      refreshHashed,
		PairID:    pairId,
		Identity:  params.Identity,
		UserAgent: params.UserAgent,
		IP:        params.IP,
		ExpiresAt: refreshExpires,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: can not save refresh token in storage", op)
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
func generateToken(payload *TokenPayload, expires time.Time) (string, error) {
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
		return "", err
	}
	claims["identity"] = identity

	// Signing with secret key
	t, err := access.SignedString([]byte(config.Get(config.EnvSecret)))
	if err != nil {
		return "", err
	}

	return t, nil
}
