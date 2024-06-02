package main

import (
	"cook/handlers"
	"cook/workers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/cook/list", handlers.HandleDeliveryListOrders)

	cook := workers.MakeRabbitMQListener()

	cook.Listen()

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:3142",
	}
	fmt.Println("Starting server")
	log.Fatalln(srv.ListenAndServe())
}
