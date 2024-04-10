package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy Learning")
}

func main() {
	// Registering the handler for the /hello endpoint
	http.HandleFunc("/", helloHandler)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
