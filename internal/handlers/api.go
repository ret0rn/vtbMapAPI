package handlers

import "github.com/ret0rn/vtbMapAPI/internal/service"

type Implementation struct {
	srv *service.Service
}

func NewPublicAPI(service *service.Service) *Implementation {
	return &Implementation{
		srv: service,
	}
}
