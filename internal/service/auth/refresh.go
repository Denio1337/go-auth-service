package auth

import (
	"app/internal/service/utils"
	"app/internal/storage"
	"app/internal/storage/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Update tokens
func Refresh(params *RefreshParams) (*Tokens, error) {
	const op = "service/refresh"

	// Get refresh token from storage
	storageRefresh, err := storage.GetRefreshTokenByPairID(params.PairID)
	if err != nil {
		return nil, fmt.Errorf("%s: can not get refresh token by pair id", op)
	}

	// Check user agent change
	if params.UserAgent != storageRefresh.UserAgent {
		return nil, ErrUserAgentChanged
	}

	// Revoke previous refresh token
	revoked, err := storage.RevokeRefreshTokenByID(storageRefresh.ID)
	if err != nil || revoked == 0 {
		return nil, fmt.Errorf("%s: can not revoke refresh token", op)
	}

	// Generate general tokens parameters
	pairId := uuid.New().String()
	payload := &TokenPayload{
		Username: params.Username,
		ID:       params.ID,
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

	// Hash refresh token
	refreshHashed, err := utils.Hash(refresh)
	if err != nil {
		return nil, fmt.Errorf("%s: can not hash refresh token", op)
	}

	// Save new refresh token
	err = storage.AddRefreshToken(&model.RefreshToken{
		UserID:    params.ID,
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

	// Return different error with tokens if ip changed
	var ipChangedError error
	if params.IP != storageRefresh.IP {
		ipChangedError = ErrIPChanged
	}

	return &Tokens{
		Access:  access,
		Refresh: refresh,
	}, ipChangedError
}
