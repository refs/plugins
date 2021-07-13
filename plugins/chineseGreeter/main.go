package main

import (
	"log"
	"net/http"
)

// HeyyHandler is just gibberish
func HeyyHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("哟?")); err != nil {
		log.Fatal(err)
	}
}

// Serve starts an HTTP server
func Serve() {
	if err := http.ListenAndServe(":8081", http.HandlerFunc(HeyyHandler)); err != nil {
		log.Fatal(err)
	}
}
