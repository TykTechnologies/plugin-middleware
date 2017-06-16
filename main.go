package main

import (
	"log"
	"net/http"
)

// SampleMiddleware is a sample middleware.
func SampleMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling SampleMiddleware")
}
