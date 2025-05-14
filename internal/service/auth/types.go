package auth

import (
	"app/internal/storage/model"
	"time"
)

type (
	// Tokens structure
	Tokens struct {
		Access  string
		Refresh string
	}

	// Login method params
	LoginParams struct {
		Username string
		Password string
		Identity string
	}

	// Login method result
	LoginResult struct {
		*Tokens
		ID       uint
		Username string
	}

	// Register method params
	RegisterParams struct {
		Username string
		Password string
	}

	// Register method result
	RegisterResult struct {
		ID       uint
		Username string
	}

	// Me method params
	MeParams struct {
		ID uint
	}

	// Me method result
	MeResult struct {
		ID       uint
		Username string
	}

	// Logout method params
	LogoutParams struct {
		Refresh  string
		Identity string
	}

	// Payload for JWT tokens
	TokenPayload struct {
		*model.User
		Identity string
		PairID   string
	}

	// Token live time
	TokenLiveTime time.Time
)

const (
	// Token type
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	TokenLiveTimeAccess  = 3 * time.Minute
	TokenLiveTimeRefresh = 3 * time.Hour
)
