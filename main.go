package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define a simple test route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is up and running!")
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
