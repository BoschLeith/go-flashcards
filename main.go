package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func study(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Study flashcards"))
}

func flashcardView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "View flashcard with ID %d", id)
}

func flashcardCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new flashcard form"))
}

func flashcardCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save new flashcard"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /study", study)
	mux.HandleFunc("GET /flashcard/view/{id}", flashcardView)
	mux.HandleFunc("GET /flashcard/create", flashcardCreate)
	mux.HandleFunc("POST /flashcard/create", flashcardCreatePost)

	log.Print("Server started on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
