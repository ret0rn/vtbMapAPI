package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

type Repo interface {
	GetHandlingList(context.Context, *model.HandlingFilter) (model.HandlingList, error)
	GetOfficeLocation(context.Context, *model.OfficeLocationFilter) (model.OfficeLocationList, error)
	GetRatesByOffice(context.Context, []int64) (map[int64]float64, error)
	GetCountPeopleByOffice(context.Context, []int64, int64) (map[int64]int64, error)
}

type Repository struct {
	MasterPool *pgxpool.Pool
	timeout    time.Duration
}

func NewRepository(masterPool *pgxpool.Pool, timeout time.Duration) Repo {
	return &Repository{
		MasterPool: masterPool,
		timeout:    timeout,
	}
}
