package service

import (
	"context"
	"fmt"

	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (s *Service) GetHandlingByFilter(ctx context.Context, filter *model.HandlingFilter) (*model.Handling, error) {
	list, err := s.GetHandlingListByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return &model.Handling{}, nil
	}
	return list[0], nil
}

func (s *Service) GetHandlingListByFilter(ctx context.Context, filter *model.HandlingFilter) (model.HandlingList, error) {
	data, ok := s.cache.Get(fmt.Sprintf(model.CacheHandlingKey, filter.ClientType, filter.HandlingId))
	if ok {
		return data.(model.HandlingList), nil
	}
	list, err := s.repo.GetHandlingList(ctx, filter)
	if err != nil {
		return nil, err
	}
	s.cache.Set(fmt.Sprintf(model.CacheHandlingKey, filter.ClientType, filter.HandlingId), list, model.DefaultTTl)
	return list, nil
}
