package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (s *Service) GetOfficeLocationList(ctx context.Context, filter *model.OfficeLocationFilter) (model.OfficeLocationList, map[int64]int64, map[int64]float64, time.Duration, error) {
	list, err := s.repo.GetOfficeLocation(ctx, filter)
	if err != nil {
		return nil, nil, nil, 0, errors.Wrapf(err, "GetOfficeLocation error")
	}
	var countPeople map[int64]int64
	var handlingDuration time.Duration
	if filter.ClientType != 0 && filter.HandlingType != 0 {
		countPeople, err = s.repo.GetCountPeopleByOffice(ctx, list.GetOfficeIds(), filter.HandlingType)
		if err != nil {
			return nil, nil, nil, 0, errors.Wrapf(err, "GetWorkloadByOffice error")
		}
		var handling *model.Handling
		handling, err = s.GetHandlingByFilter(ctx, &model.HandlingFilter{HandlingId: filter.HandlingType, ClientType: filter.ClientType}) // TODO: сделать подбор услуги по офису
		if err != nil {
			return nil, nil, nil, 0, errors.Wrapf(err, "GetWorkloadByOffice error")
		}
		handlingDuration = *handling.HandlingDuration
	}
	ratesMap, err := s.repo.GetRatesByOffice(ctx, list.GetOfficeIds())
	if err != nil {
		return nil, nil, nil, 0, errors.Wrapf(err, "GetRatesByOffice error")
	}

	return list, countPeople, ratesMap, handlingDuration, nil
}
