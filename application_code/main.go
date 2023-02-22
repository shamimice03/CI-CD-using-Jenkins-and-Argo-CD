package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the function that handles incoming requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	// Register the handler function to handle all requests to the root URL path
	http.HandleFunc("/", handler)

	// Start the HTTP server and listen for incoming requests on port 8080
	http.ListenAndServe(":8080", nil)
}
