package main

import (
	"log"
	"net/http"

	"product-search-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", handlers.SearchProducts).Methods("GET")
	r.HandleFunc("/suggestions", handlers.SuggestProducts).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
