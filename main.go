package main

import (
	"log"
	"main/validator"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	cardValidator := validator.NewCardValidator()
	validateHandler := validator.NewValidateHandler(cardValidator)
	mux.Handle("POST /validate", validateHandler)

	log.Println("starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
