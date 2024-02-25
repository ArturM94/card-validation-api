package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}