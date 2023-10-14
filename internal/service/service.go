package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ret0rn/vtbMapAPI/internal/config"
	"github.com/ret0rn/vtbMapAPI/internal/repository"
	"github.com/ret0rn/vtbMapAPI/utils/cache"
)

type Service struct {
	repo  repository.Repo
	cache *cache.Cache
}

func NewService(ctx context.Context) (*Service, error) {
	const defaultTimeout = 5 * time.Second // дефолтный таймаут на запрос

	masterPool, err := pgxpool.Connect(ctx, config.GetMasterPool())
	if err != nil {
		return nil, err
	}
	var repo = repository.NewRepository(masterPool, defaultTimeout)
	return &Service{
		repo:  repo,
		cache: cache.NewCache(),
	}, nil
}
