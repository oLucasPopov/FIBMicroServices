package Bff

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:3143",
	}

	log.Fatalln(srv.ListenAndServe())
}
