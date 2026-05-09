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

type OrderMessage struct {
	UserID   uint   `json:"user_id"`
	OrderID  uint   `json:"order_id"`
	ItemName string `json:"item_name"`
	Nums     int    `json:"nums"`
}

func NewOrderService(db OrderRepository, publisher KafkaPublisher) *OrderService {
	return &OrderService{db: db, publisher: publisher}
}

func (s *OrderService) CreateOrder(itemName string, nums int) error {
	order := &model.Order{
		ItemName: itemName,
		Nums:     nums,
	}
	savedOrder, err := s.db.CreateOrder(order)
	if err != nil {
		return err
	}
	orderMessage := OrderMessage{
		UserID:   savedOrder.UserID,
		OrderID:  savedOrder.ID,
		ItemName: savedOrder.ItemName,
		Nums:     savedOrder.Nums,
	}
	orderBytes, _ := json.Marshal(orderMessage)
	if err := s.publisher.Publish("orders-created", []byte(fmt.Sprintf("%d", savedOrder.ID)), orderBytes); err != nil {
		return err
	}
	return nil
}
