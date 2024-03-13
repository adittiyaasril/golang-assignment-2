package repo

import (
    "gorm.io/gorm"
    "assignment-2/core"
)

type OrderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(order *core.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetOrders() ([]core.Order, error) {
	var orders []core.Order
	if err := r.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrder(order *core.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&core.Order{}, id).Error
}
