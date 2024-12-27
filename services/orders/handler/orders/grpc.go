package handler

import (
	"context"
	"github.com/spayder/kitchen-api/services/common/genproto/orders"
	"github.com/spayder/kitchen-api/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(grpc *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrderGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    23,
		CustomerID: 232,
		ProductID:  123,
		Quantity:   1,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
