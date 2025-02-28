package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /flashcard/view/{id}", flashcardView)
	mux.HandleFunc("GET /flashcard/create", flashcardCreate)
	mux.HandleFunc("POST /flashcard/create", flashcardCreatePost)

	log.Print("starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
