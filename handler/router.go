package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ConfigureRouter(handler InfoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcards", handler.PostFlashcardHandler).Methods("POST")
	r.HandleFunc("/flashcards", handler.GetFlashcardsHandler).Methods("GET")

	return r
}

func NewServer(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
