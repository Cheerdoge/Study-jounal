package config

import (
	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(brokers []string, clientID string) (*KafkaPublisher, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V4_2_0_0
	config.ClientID = clientID
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaPublisher{producer: producer}, nil
}

func (kp *KafkaPublisher) Publish(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	_, _, err := kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func (kp *KafkaPublisher) Close() error {
	return kp.producer.Close()
}
