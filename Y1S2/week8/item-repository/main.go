package main

import (
	"context"
	"item-repository/config"
	"item-repository/handler"
	"item-repository/registry"
	"item-repository/repository"
	"item-repository/router"
	"item-repository/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
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
	handler := handler.NewInventoryHandler(svc)

	r := gin.Default()
	router.RegisterRoutes(r, handler)
	log.Println("Starting HTTP server on :8090")
	go func() {
		if err := r.Run(":8090"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

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
			if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
				log.Printf("Error consuming messages: %v", err)
			}
			if ctx.Err() != nil {
				log.Printf("Kafka consumer context error: %v", ctx.Err())
				return
			}
		}
	}()
	reg, err := registry.NewServiceRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}

	err = reg.Register("order-create-service", "127.0.0.1:8080", 5)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	log.Println("Shutting down inventory service...")
	if err := reg.Deregister(); err != nil {
		log.Printf("Failed to deregister service: %v", err)
	}
	cancel()
}
