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

	CookieParams struct {
		Name    string
		Value   string
		Expires time.Time
	}
)

const (
	CookieNameAccessToken  = "auth_access"
	CookieNameRefreshToken = "auth_refresh"
)
