package service

import (
	"assignment-2/core"
	"assignment-2/repo"
)

type OrderService struct {
	orderRepo *repo.OrderRepository
}

func NewOrderService(orderRepo *repo.OrderRepository) *OrderService {
	return &OrderService{orderRepo}
}

func (s *OrderService) CreateOrder(order *core.Order) error {
	return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrders() ([]core.Order, error) {
	return s.orderRepo.GetOrders()
}

func (s *OrderService) UpdateOrder(order *core.Order) error {
	return s.orderRepo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.orderRepo.DeleteOrder(id)
}
