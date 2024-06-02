package handlers

import (
	"delivery/controller"
	"delivery/service"
	"net/http"
)

func HandleDeliveryListOrders(w http.ResponseWriter, r *http.Request) {
	dc := controller.MakeDeliveryController(&service.OrderService{}, &service.DeliveryService{}, service.MakeRabbitMQSender())
	dc.ListOrders(w)
}
