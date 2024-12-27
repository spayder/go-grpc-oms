package main

import (
	handler "github.com/spayder/kitchen-api/services/orders/handler/orders"
	"github.com/spayder/kitchen-api/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	r := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrderHttpHandler(orderService)
	orderHandler.RegisterRoute(r)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, r)
}
