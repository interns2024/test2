package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Modified response to include path information
	fmt.Fprintf(w, "Happy Learning on path %s!", r.URL.Path)
}

func main() {
	// Registering the handler for the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
