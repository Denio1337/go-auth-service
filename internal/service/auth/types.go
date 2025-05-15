package auth

import (
	"errors"
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
		Username  string
		Password  string
		Identity  string
		IP        string
		UserAgent string
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

	// Refresh method params
	RefreshParams struct {
		Username  string
		ID        uint
		Identity  string
		PairID    string
		UserAgent string
		IP        string
	}

	// Payload for JWT tokens
	TokenPayload struct {
		Username string
		ID       uint
		Identity string
		PairID   string
	}
)

// Errors
var (
	ErrUserExists = errors.New("service/register: user already exists")

	ErrUserAgentChanged = errors.New("service/refresh: user agent changed")
	ErrIPChanged        = errors.New("service/refresh: ip changed")
)

const (
	// Token type
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	// Token live time
	TokenLiveTimeAccess  = 3 * time.Minute
	TokenLiveTimeRefresh = 3 * time.Hour
)
