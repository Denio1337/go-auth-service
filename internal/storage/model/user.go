package model

import "gorm.io/gorm"

// User structure
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:50;" validate:"required,min=3,max=50" json:"username"`
	Password string `gorm:"not null;size:50" validate:"required,min=8,max=50" json:"password"`
}
