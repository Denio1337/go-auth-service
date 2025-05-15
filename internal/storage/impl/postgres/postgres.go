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
	p := config.Get(config.EnvDBPort)
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	// Define data source name
	dsn := fmt.Sprintf(
		"host=db port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Get(config.EnvDBUser),
		config.Get(config.EnvDBPassword),
		config.Get(config.EnvDBName),
	)

	// Try to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schemas to database
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.RefreshToken{})

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
		if isUniqueConstraintError(err) {
			return contract.ErrDuplicatedKey
		}

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

func (s *PostgresStorage) RevokeRefreshTokenByPairID(pairID string) (uint, error) {
	result := s.db.Delete(&model.RefreshToken{}, "pair_id = ?", pairID)

	if err := result.Error; err != nil {
		return 0, err
	}

	return uint(result.RowsAffected), nil
}

func (s *PostgresStorage) RevokeRefreshTokenByIdentity(identity string) (uint, error) {
	result := s.db.Unscoped().Delete(&model.RefreshToken{}, "identity = ?", identity)

	if err := result.Error; err != nil {
		return 0, err
	}

	return uint(result.RowsAffected), nil
}

func (s *PostgresStorage) RevokeRefreshTokenByID(id uint) (uint, error) {
	result := s.db.Unscoped().Delete(&model.RefreshToken{}, id)

	if err := result.Error; err != nil {
		return 0, err
	}

	return uint(result.RowsAffected), nil
}
