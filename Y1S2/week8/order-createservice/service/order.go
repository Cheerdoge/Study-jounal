package service

import "order-payment-kafka/order-createservice/model"

type KafkaPublisher interface {
	Publish(topic string, key, value []byte) error
	Close() error
}

type OrderService struct {
	publisher KafkaPublisher
}

func NewOrderService(publisher KafkaPublisher) *OrderService {
	return &OrderService{publisher: publisher}
}

func (s *OrderService) CreateOrder(item string) error {
	order := &model.Order{
		Item:   item,
		Status: "created",
	}
	return s.publisher.Publish("orders", nil, []byte(order.Item))
}
