package service

import (
	"encoding/json"
	"fmt"
	"order-payment-kafka/order-createservice/model"
)

type KafkaPublisher interface {
	Publish(topic string, key, value []byte) error
	Close() error
}

type OrderRepository interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	GetOrderByID(id uint) (*model.Order, error)
}

type OrderService struct {
	db        OrderRepository
	publisher KafkaPublisher
}

func NewOrderService(db OrderRepository, publisher KafkaPublisher) *OrderService {
	return &OrderService{db: db, publisher: publisher}
}

func (s *OrderService) CreateOrder(item string) error {
	order := &model.Order{
		Item:   item,
		Status: "created",
	}
	savedOrder, err := s.db.CreateOrder(order)
	if err != nil {
		return err
	}
	orderBytes, _ := json.Marshal(savedOrder)
	if err := s.publisher.Publish("orders-created", []byte(fmt.Sprintf("%d", savedOrder.ID)), orderBytes); err != nil {
		return err
	}
	return nil
}
