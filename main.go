package main

import (
	"golang-webapi/api"
	"net/http"
)

func main() {
	// This is only for local testing
	http.HandleFunc("/", api.Handler)

	// Start server (For local development only)
	http.ListenAndServe(":8080", nil)
}
