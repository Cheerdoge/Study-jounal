package main

import (
	"context"
	"fmt"
	"log"
	"week12/hello_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		msg := fmt.Sprintf("Failed to connect: %v", err)
		log.Fatalf(msg)
	}
	defer conn.Close()
	// 初始化客户端
	client := hello_grpc.NewHelloServiceClient(conn)
	result, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "2222",
		Message: "ok",
	})
	fmt.Println(result, err)
}
