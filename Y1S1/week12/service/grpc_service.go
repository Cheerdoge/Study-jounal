package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"week12/hello_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HelloService struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (HelloService) SayHello(ctx context.Context, req *hello_grpc.HelloRequest) (resp *hello_grpc.HelloResponse, err error) {
	fmt.Println("收到请求:", req.GetName())
	resp = &hello_grpc.HelloResponse{
		Name:    req.GetName(),
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}
	return resp, nil
}

func main() {
	//监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//创建gRPC服务器
	s := grpc.NewServer()
	//注册gRPC服务
	service := HelloService{}
	hello_grpc.RegisterHelloServiceServer(s, service)
	fmt.Println("服务已启动")
	//反射服务
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
