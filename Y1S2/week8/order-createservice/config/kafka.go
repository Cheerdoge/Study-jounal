package config

import (
	"log"

	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

var Publisher *KafkaPublisher

func NewKafkaPublisher() error {
	config := sarama.NewConfig()
	config.Version = sarama.V4_2_0_0
	config.ClientID = "order-service-producer"
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		return err
	}

	Publisher = &KafkaPublisher{producer: producer}
	return nil
}

func (kp *KafkaPublisher) Publish(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	partition, offset, err := kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("Message sent to topic %s, partition %d at offset %d", topic, partition, offset)
	return nil
}

func (kp *KafkaPublisher) Close() error {
	return kp.producer.Close()
}
