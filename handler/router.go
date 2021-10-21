package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter(handler InfoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcards", handler.PostNewDeck).Methods("POST")
	r.HandleFunc("/flashcards", handler.PostNewInfo).Methods("POST")

	return r
}

func NewServer(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
