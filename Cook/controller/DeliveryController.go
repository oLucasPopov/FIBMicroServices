package controller

import (
	"cook/service"
	"cook/types"
	"encoding/json"
	"net/http"
)

type DeliveryController struct {
	orderService *service.OrderService
}

func MakeDeliveryController(orderService *service.OrderService) *DeliveryController {
	return &DeliveryController{
		orderService: orderService,
	}
}

func (dc *DeliveryController) ListOrders(w http.ResponseWriter) {
	orders, err := dc.orderService.ListOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&types.Error{Message: err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(orders)
}
