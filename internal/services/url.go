package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"
	"url-shortener/utils"
)

var validate = validator.New()

type Service interface {
	CreateShortURL(ctx context.Context, full string, expiryHours int) (*models.ShortURL, error)
	ResolveURL(ctx context.Context, code string) (*models.ShortURL, error)
	GetBaseURL() string
}

type urlService struct {
	store   repositories.Store
	cache   repositories.Cache
	baseURL string
}

func (s *urlService) GetBaseURL() string {
	return s.baseURL
}

func NewService(store repositories.Store, cache repositories.Cache, baseURL string) Service {
	return &urlService{
		store:   store,
		cache:   cache,
		baseURL: baseURL,
	}
}

func (s *urlService) CreateShortURL(ctx context.Context, full string, expiryHours int) (*models.ShortURL, error) {
	if err := validate.Var(full, "required,url"); err != nil {
		return nil, errors.New("invalid URL")
	}
	code, err := utils.GenerateCode(6)
	if err != nil {
		return nil, err
	}
	expiry := time.Duration(expiryHours) * time.Hour
	short := &models.ShortURL{Code: code, FullURL: full, CreatedAt: time.Now(), Expiry: expiry}
	if err := s.store.Save(ctx, short); err != nil {
		return nil, err
	}
	return short, nil
}

func (s *urlService) ResolveURL(ctx context.Context, code string) (*models.ShortURL, error) {
	cached, err := s.cache.Get(ctx, code)
	if err == nil && cached != nil {
		_ = s.store.IncrementClicks(ctx, code)
		return cached, nil
	}

	u, err := s.store.Find(ctx, code)
	if err != nil {
		return nil, err
	}

	if time.Since(u.CreatedAt) > u.Expiry {
		return nil, errors.New("url expired")
	}

	_ = s.cache.Set(ctx, u, time.Until(u.CreatedAt.Add(u.Expiry)))

	_ = s.store.IncrementClicks(ctx, code)

	return u, nil
}
