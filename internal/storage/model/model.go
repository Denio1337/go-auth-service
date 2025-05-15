package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	// User structure
	User struct {
		gorm.Model

		// General information
		Username string `gorm:"uniqueIndex;not null;size:50;" validate:"required,min=3,max=50" json:"username"`
		Password string `gorm:"not null;size:60" validate:"required,min=8,max=50" json:"password"`

		// Refresh token foreign key
		RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
	}

	// Refresh token structure
	RefreshToken struct {
		gorm.Model

		// User reference
		UserID uint `gorm:"not null"`
		User   User `gorm:"foreignKey:UserID"`

		// Hash of refresh token
		Hash string `gorm:"not null"`

		// Pair ID for access/refresh tokens
		PairID string `gorm:"not null"`

		// Refresh token info
		Identity  string    `gorm:"not null"`
		UserAgent string    `gorm:"not null"`
		IP        string    `gorm:"not null"`
		ExpiresAt time.Time `gorm:"not null"`
	}
)
