package main

import (
	"context"
	"item-repository/config"
	"item-repository/handler"
	"item-repository/repository"
	"item-repository/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadConfig()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	defer config.CloseDatabase(db)

	repo := repository.NewInventoryRepository(db)
	svc := service.NewItemService(repo)
	kafkaHandler := handler.NewKafkaHandler(svc)

	consumerGroup, err := config.NewKafkaConsumer()
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
		return
	}
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	topics := []string{"orders-created"}

	go func() {
		for {
			log.Printf("Starting Kafka consumer for topics: %v", topics)
			if err := consumerGroup.Consume(ctx, topics, kafkaHandler); err != nil {
				log.Printf("Error consuming messages: %v", err)
			}
			if ctx.Err() != nil {
				log.Printf("Kafka consumer context error: %v", ctx.Err())
				return
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	log.Println("Shutting down inventory service...")
	cancel()
}
