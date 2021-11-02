package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gpark1005/FlashCardsTeamOne/cards"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
)

type Service interface {
	PostNewMatching(card cards.Matching) error
	PostNewMultipleChoice(card cards.MultipleChoice) error
	PostNewInfo(card cards.Info) error
	PostNewQNA(card cards.QNA) error
	PostNewTORF(card cards.TrueOrFalse) error
	GetAllFlashcards() (repo.Db, error)
	GetByType(input string) (repo.DbType, error)
	DeleteById(input string) error
	UpdateById(input string, card map[string]interface{}) error
	GetByCategory(input string) (repo.DbType, error)
	GetById(input string) (repo.DbType, error)
}

type CardHandler struct {
	Svc Service
}

func NewCardHandler(s Service) CardHandler {
	return CardHandler{
		Svc: s,
	}
}

var CardType map[string]interface{}

func (ch CardHandler) PostFlashcardHandler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusBadRequest)
		return
	}

	if cType, ok := CardType["type"]; ok {
		switch cType {
		case "matching":
			matchCard := cards.Matching{}
			err = json.Unmarshal(data, &matchCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			err = ch.Svc.PostNewMatching(matchCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "multiple":
			multipleCard := cards.MultipleChoice{}
			err = json.Unmarshal(data, &multipleCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			err = ch.Svc.PostNewMultipleChoice(multipleCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "info":
			infoCard := cards.Info{}
			err = json.Unmarshal(data, &infoCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			err = ch.Svc.PostNewInfo(infoCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "qAndA":
			qAndACard := cards.QNA{}
			err = json.Unmarshal(data, &qAndACard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			err = ch.Svc.PostNewQNA(qAndACard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "tOrF":
			torfCard := cards.TrueOrFalse{}
			err = json.Unmarshal(data, &torfCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			err = ch.Svc.PostNewTORF(torfCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "invaild type", http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (ch CardHandler) GetFlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	myDb, err := ch.Svc.GetAllFlashcards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := json.MarshalIndent(myDb, "", " ")
	if err != nil {
		http.Error(w, "unable to encode database", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(db)
}

func (ch CardHandler) GetByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["type"]

	getType, err := ch.Svc.GetByType(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	flashcard, err := json.MarshalIndent(getType, "", "	")
	if err != nil {
		http.Error(w, "unable to encode database", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(flashcard)
}

func (ch CardHandler) DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := ch.Svc.DeleteById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (ch CardHandler) UpdateByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to access database", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusBadRequest)
		return
	}

	err = ch.Svc.UpdateById(id, CardType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (ch CardHandler) GetByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	getCategory, err := ch.Svc.GetByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	flashcard, err := json.MarshalIndent(getCategory, "", "	")
	if err != nil {
		http.Error(w, "unable to encode database", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(flashcard)
}

func (ch CardHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	getId, err := ch.Svc.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	flashcard, err := json.MarshalIndent(getId, "", "	")
	if err != nil {
		http.Error(w, "unable to encode database", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(flashcard)
}
