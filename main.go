package main

import (
	"log"
	"net/http"
	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")

	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
