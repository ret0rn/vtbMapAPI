package service

import (
	"context"

	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (s *Service) GetHandlingListByFilter(ctx context.Context, filter *model.HandlingFilter) ([]*model.Handling, error) {
	return s.repo.GetHandlingList(ctx, filter)
}
