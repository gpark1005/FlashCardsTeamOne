package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter(handler InfoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcards", handler.PostFlashcardHandler).Methods("POST")
	r.HandleFunc("/flashcards", handler.GetFlashcardsHandler).Methods("GET")
	r.HandleFunc("/flashcards/{type}", handler.GetByTypeHandler).Methods("GET")
	r.HandleFunc("/flashcards/{id}", handler.DeleteByIdHandler).Methods("DELETE")
	r.HandleFunc("/flashcards/{id}", handler.UpdateByIdHandler).Methods("PUT")

	return r
}

func NewServer(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
