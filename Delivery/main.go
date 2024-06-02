package main

import (
	"delivery/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/order/new", handlers.HandleDeliveryNewOrder)
	r.HandleFunc("/order/list", handlers.HandleDeliveryListOrders)
	r.HandleFunc("/menu", handlers.HandleDeliveryListMenu)

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:3141",
	}

	log.Fatalln(srv.ListenAndServe())
}
