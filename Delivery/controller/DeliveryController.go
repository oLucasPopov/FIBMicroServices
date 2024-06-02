package controller

import (
	"delivery/service"
	"delivery/types"
	"encoding/json"
	"io"
	"net/http"
)

type DeliveryController struct {
	deliveryService *service.DeliveryService
	orderService    *service.OrderService
	rabbitMQService *service.RabbitMQSender
}

func MakeDeliveryController(
	orderService *service.OrderService,
	deliveryService *service.DeliveryService,
	rabbitMQService *service.RabbitMQSender,
) *DeliveryController {
	return &DeliveryController{
		orderService:    orderService,
		deliveryService: deliveryService,
		rabbitMQService: rabbitMQService,
	}
}

func (dc *DeliveryController) NewOrder(order io.ReadCloser, w http.ResponseWriter) {
	defer dc.rabbitMQService.CloseConnection()

	var newOrder types.Order
	err := json.NewDecoder(order).Decode(&newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&types.Error{
			Message: err.Error(),
		})

		return
	}

	resOrder, err := dc.orderService.NewOrder(newOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&types.Error{
			Message: err.Error(),
		})

		return
	}

	if err = dc.rabbitMQService.ConvertAndSend(resOrder); err != nil {
		_ = json.NewEncoder(w).Encode(&types.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resOrder)
}

func (dc *DeliveryController) ListOrders(w http.ResponseWriter) {
	orders, err := dc.orderService.ListOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&types.Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(orders)
}

func (dc *DeliveryController) ListMenu(w http.ResponseWriter) {
	menu, err := dc.deliveryService.ListMenu()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&types.Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(menu)
}
