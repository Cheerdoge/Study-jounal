package handler

import (
	"encoding/json"
	"item-repository/model"
	"log"

	"github.com/IBM/sarama"
)

type InventoryService interface {
	GetItemByName(name string) (model.Item, error)
	GetItemByID(id uint) (model.Item, error)
	GetItems() ([]model.Item, error)
	ProcessOrder(orderID uint, itemname string, nums int) error
}

type KafkaHandler struct {
	Service InventoryService
}

type OrderMessage struct {
	OrderID  uint   `json:"order_id"`
	ItemName string `json:"item_name"`
	Nums     int    `json:"nums"`
}

func NewKafkaHandler(service InventoryService) *KafkaHandler {
	return &KafkaHandler{
		Service: service,
	}
}

func (h *KafkaHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: topic = %s, partition = %d, offset = %d", message.Topic, message.Partition, message.Offset)
		var orderMsg OrderMessage
		if err := json.Unmarshal(message.Value, &orderMsg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			session.MarkMessage(message, "")
			continue
		}
		log.Printf("Processing order: ID = %d, Item = %s, Nums = %d", orderMsg.OrderID, orderMsg.ItemName, orderMsg.Nums)
		if err := h.Service.ProcessOrder(orderMsg.OrderID, orderMsg.ItemName, orderMsg.Nums); err != nil {
			log.Printf("Error processing order: %v", err)
			session.MarkMessage(message, "")
			continue
		}
		session.MarkMessage(message, "")
	}
	return nil
}
