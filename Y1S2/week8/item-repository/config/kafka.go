package config

import (
	"log"

	"github.com/IBM/sarama"
)

func NewKafkaConsumer() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V4_2_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	brokers := []string{"127.0.0.1:9092"}
	groupId := "inventory-service-group"
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		log.Printf("Error creating consumer group: %v", err)
		return nil, err
	}

	return consumerGroup, nil
}
