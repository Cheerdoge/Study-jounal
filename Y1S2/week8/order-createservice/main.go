package main

import (
	"log"
	"order-payment-kafka/order-createservice/config"
	"order-payment-kafka/order-createservice/handler"
	"order-payment-kafka/order-createservice/registry"
	"order-payment-kafka/order-createservice/repository"
	"order-payment-kafka/order-createservice/router"
	"order-payment-kafka/order-createservice/service"
	"os"
	"os/signal"
	"syscall"

	pb "order-payment-kafka/order-createservice/pb"

	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.LoadConfig()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDatabase(db)

	if err := config.NewKafkaPublisher(); err != nil {
		log.Fatalf("Failed to create Kafka publisher: %v", err)
	}
	defer config.Publisher.Close()

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		log.Fatalf("Failed to create etcd resolver: %v", err)
	}
	target := "etcd:////services/inventory-grpc-service"

	conn, err := grpc.NewClient(target,
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)

	if err != nil {
		log.Fatalf("Failed to connect to inventory gRPC service: %v", err)
	}
	defer conn.Close()

	invClient := pb.NewInventoryServiceClient(conn)

	orderrepository := repository.NewOrderRepository(db)
	orderservice := service.NewOrderService(orderrepository, config.Publisher, invClient)
	orderhandler := handler.NewOrderHandler(orderservice)

	go func() {
		r := gin.Default()
		router.RegisterRoutes(r, orderhandler)
		log.Println("Starting server on :8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
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
}
