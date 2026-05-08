package service

import "order-payment-kafka/order-createservice/model"

// type KafkaPublisher interface {
// 	Publish(topic string, key, value []byte) error
// 	Close() error
// }

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetOrderByID(id uint) (*model.Order, error)
}

type OrderService struct {
	db OrderRepository
	// publisher KafkaPublisher
}

func NewOrderService(db OrderRepository) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(item string) error {
	order := &model.Order{
		Item:   item,
		Status: "created",
	}
	return s.db.CreateOrder(order)
}
