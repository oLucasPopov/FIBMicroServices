package handlers

import (
	"delivery/controller"
	"delivery/service"
	"net/http"
)

func HandleDeliveryListMenu(w http.ResponseWriter, r *http.Request) {
	dc := controller.MakeDeliveryController(
		&service.OrderService{},
		&service.DeliveryService{},
		service.MakeRabbitMQSender(),
	)

	dc.ListMenu(w)
}
