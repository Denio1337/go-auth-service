package postgres

import (
	"app/internal/config"
	"app/internal/storage/contract"
	"app/internal/storage/model"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage implementation structure
type PostgresStorage struct {
	db *gorm.DB
}

// Create new PostgreSQL storage implementation
func New() (contract.Storage, error) {
	// Parse port from environment
	p := config.Get("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	// Define data source name
	dsn := fmt.Sprintf(
		"host=db port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
	)

	// Trying to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection opened")

	// Migrate schemas to database
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.RefreshToken{})
	fmt.Println("Database migrated")

	return &PostgresStorage{db: db}, nil
}

// Interface

func (s *PostgresStorage) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (s *PostgresStorage) AddUser(user *model.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := s.db.Where("ID = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (s *PostgresStorage) GetRefreshTokenByPairID(pairID string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	if err := s.db.Where(&model.RefreshToken{PairID: pairID}).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &token, nil
}

func (s *PostgresStorage) GetRefreshTokenByIdentity(identity string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	if err := s.db.Where(&model.RefreshToken{Identity: identity}).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &token, nil
}

func (s *PostgresStorage) AddRefreshToken(token *model.RefreshToken) error {
	if err := s.db.Create(token).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) RevokeRefreshTokenByPairID(pairID string) error {
	if err := s.db.Delete(&model.RefreshToken{}, "pair_id = ?", pairID).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) RevokeRefreshTokenByIdentity(identity string) error {
	if err := s.db.Unscoped().Delete(&model.RefreshToken{}, "identity = ?", identity).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) RevokeRefreshTokenByID(id uint) error {
	if err := s.db.Unscoped().Delete(&model.RefreshToken{}, id).Error; err != nil {
		return err
	}

	return nil
}
