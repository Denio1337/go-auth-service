package auth

import "time"

type (
	// DTO for GET /auth/login
	LoginDTO struct {
		Username string `json:"username" validate:"required,min=5,max=20"`
		Password string `json:"password" validate:"required,min=8,max=30"`
	}

	// Response for GET /auth/login
	LoginResponse struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	// DTO for GET /auth/register
	RegisterDTO struct {
		Username string `json:"username" validate:"required,min=5,max=20"`
		Password string `json:"password" validate:"required,min=8,max=30"`
	}

	// Response for GET /auth/register
	RegisterResponse struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	// Response for GET /auth/me
	MeResponse struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	// Parameters to set/clear cookies
	CookieParams struct {
		Name    string
		Value   string
		Expires time.Time
	}

	WebhookPayload struct {
		ID        uint
		Timestamp time.Time
		Message   string
	}
)

const (
	// Auth cookie names
	CookieNameAccessToken  = "auth_access"
	CookieNameRefreshToken = "auth_refresh"
)
