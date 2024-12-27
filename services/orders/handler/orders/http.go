package handler

import (
	"github.com/spayder/kitchen-api/services/common/genproto/orders"
	"github.com/spayder/kitchen-api/services/common/util"
	"github.com/spayder/kitchen-api/services/orders/types"
	"net/http"
)

type OrderHttpHandler struct {
	orderService types.OrderService
}

func NewOrderHttpHandler(orderService types.OrderService) *OrderHttpHandler {
	return &OrderHttpHandler{
		orderService: orderService,
	}
}

func (h *OrderHttpHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request orders.CreateOrderRequest
	err := util.ParseJSON(r, &request)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: request.GetCustomerID(),
		ProductID:  request.GetProductID(),
		Quantity:   request.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := &orders.CreateOrderResponse{
		Status: "success",
	}
	err = util.WriteJSON(w, http.StatusCreated, response)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
	}
}
