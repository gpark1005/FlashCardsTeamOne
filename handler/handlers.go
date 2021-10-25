package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
)

type Service interface {
	PostNewMatching(card cards.Matching) error
	PostNewMultiple(card cards.MultipleChoice) error
	PostNewInfo(card cards.Info) error
	PostNewQNA(card cards.QNA) error
	PostNewTORF(card cards.TrueOrFalse) error
	GetAllFlashcards() (repo.Db, error)
}

type InfoHandler struct {
	Svc Service
}

func NewInfoHandler(s Service) InfoHandler {
	return InfoHandler{
		Svc: s,
	}
}

var CardType map[string]interface{}

func (ih InfoHandler) PostFlashcardHandler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if cType, ok := CardType["type"]; ok {
		switch cType {
		case "matching":
			matchCard := cards.Matching{}
			err = json.Unmarshal(data, &matchCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = ih.Svc.PostNewMatching(matchCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "multiple":
			multipleCard := cards.MultipleChoice{}
			err = json.Unmarshal(data, &multipleCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = ih.Svc.PostNewMultiple(multipleCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "info":
			card := cards.Info{}
			err = json.Unmarshal(data, &card)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = ih.Svc.PostNewInfo(card)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "qanda":
			qandaCard := cards.QNA{}
			err = json.Unmarshal(data, &qandaCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = ih.Svc.PostNewQNA(qandaCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "torf":
			torfCard := cards.TrueOrFalse{}
			err = json.Unmarshal(data, &torfCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = ih.Svc.PostNewTORF(torfCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (ih InfoHandler) GetFlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	myDb, err := ih.Svc.GetAllFlashcards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := json.MarshalIndent(myDb, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(db)
}
