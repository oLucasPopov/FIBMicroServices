package handlers

import (
	"cook/controller"
	"cook/service"
	"net/http"
)

func HandleDeliveryListOrders(w http.ResponseWriter, r *http.Request) {
	dc := controller.MakeDeliveryController(&service.OrderService{})
	dc.ListOrders(w)
}
