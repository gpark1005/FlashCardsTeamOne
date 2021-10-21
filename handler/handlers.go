package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)

type Service interface {
}

type InfoHandler struct {
	Svc Service
}

func NewInfoHandler(s Service) InfoHandler {
	return InfoHandler{
		Svc: s,
	}
}

func (ih InfoHandler) PostNewDeck(w http.ResponseWriter, r *http.Request) {

	deck := incomingdata.Deck{}

	err := json.NewDecoder(r.Body).Decode(&deck)
	if err != nil {
		log.Fatal(err)
	}

}

func (ih InfoHandler) PostNewInfo(w http.ResponseWriter, r *http.Request) {

	card := incomingdata.Info{}

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		log.Fatal(err)
	}

	
}
