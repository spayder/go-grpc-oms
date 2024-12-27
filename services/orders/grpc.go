package main

import (
	handler "github.com/spayder/kitchen-api/services/orders/handler/orders"
	"github.com/spayder/kitchen-api/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	orderServie := service.NewOrderService()
	handler.NewOrderGrpcHandler(grpcServer, orderServie)
	log.Println("Starting gRPC server on port", s.addr)
	return grpcServer.Serve(lis)
}
