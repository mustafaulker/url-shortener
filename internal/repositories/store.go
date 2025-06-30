package repositories

import (
	"context"
	"gorm.io/gorm"
	"url-shortener/internal/models"
)

type Store interface {
	Save(ctx context.Context, url *models.ShortURL) error
	Find(ctx context.Context, code string) (*models.ShortURL, error)
	IncrementClicks(ctx context.Context, code string) error
}

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(db *gorm.DB) Store {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) Save(ctx context.Context, url *models.ShortURL) error {
	return s.db.WithContext(ctx).Create(url).Error
}

func (s *PostgresStore) Find(ctx context.Context, code string) (*models.ShortURL, error) {
	var url models.ShortURL
	err := s.db.WithContext(ctx).First(&url, "code = ?", code).Error
	return &url, err
}

func (s *PostgresStore) IncrementClicks(ctx context.Context, code string) error {
	return s.db.WithContext(ctx).
		Model(&models.ShortURL{}).
		Where("code = ?", code).
		UpdateColumn("clicks", gorm.Expr("clicks + ?", 1)).
		Error
}
