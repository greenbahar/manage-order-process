package order_service

import (
	"github.com/greenbahar/manage-order-process/order"
)

type Service interface {
	SaveOrder(order *order.Order) error
}

type Repository interface {
	SetOrder(order *order.Order) error
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) SaveOrder(order *order.Order) error {
	return s.repo.SetOrder(order)
}
