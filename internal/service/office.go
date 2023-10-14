package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (s *Service) GetOfficeLocationList(ctx context.Context, filter *model.OfficeLocationFilter) (model.OfficeLocationList, map[int64]int64, map[int64]float64, error) {
	list, err := s.repo.GetOfficeLocation(ctx, filter)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "GetOfficeLocation error")
	}
	var countPeople map[int64]int64
	if filter.ClientType != 0 && filter.HandlingType != 0 {
		countPeople, err = s.repo.GetCountPeopleByOffice(ctx, list.GetOfficeIds(), filter.HandlingType)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "GetWorkloadByOffice error")
		}
	}
	ratesMap, err := s.repo.GetRatesByOffice(ctx, list.GetOfficeIds())
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "GetRatesByOffice error")
	}

	return list, countPeople, ratesMap, nil
}
