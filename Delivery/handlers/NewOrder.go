package handlers

import (
	"delivery/controller"
	"delivery/service"
	"net/http"
)

func HandleDeliveryNewOrder(w http.ResponseWriter, r *http.Request) {
	dc := controller.MakeDeliveryController(&service.OrderService{}, &service.DeliveryService{}, service.MakeRabbitMQSender())
	dc.NewOrder(r.Body, w)
}
