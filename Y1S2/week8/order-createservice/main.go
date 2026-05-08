package main

import (
	"log"
	"order-payment-kafka/order-createservice/config"
	"order-payment-kafka/order-createservice/handler"
	"order-payment-kafka/order-createservice/repository"
	"order-payment-kafka/order-createservice/router"
	"order-payment-kafka/order-createservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDatabase(db)

	orderrepository := repository.NewOrderRepository(db)
	orderservice := service.NewOrderService(orderrepository)
	orderhandler := handler.NewOrderHandler(orderservice)
	r := gin.Default()
	router.RegisterRoutes(r, orderhandler)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
