package repository

import (
	"order-payment-kafka/order-createservice/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}
