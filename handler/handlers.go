package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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
	GetByType(input string) (repo.DbType, error)
	DeletebyId(input string) error
	UpdatebyId(input string, card map[string]interface{}) error
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
		http.Error(w,"unable to access database" , http.StatusUnprocessableEntity)
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusUnprocessableEntity)
	}

	if cType, ok := CardType["type"]; ok {
		switch cType {
		case "matching":
			matchCard := cards.Matching{}
			err = json.Unmarshal(data, &matchCard)
			if err != nil {
				http.Error(w,"unable to decode database", http.StatusUnprocessableEntity)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err := validate.Struct(matchCard)
			if err != nil {
				http.Error(w,"all fields require data", http.StatusBadRequest)
			}

			err = ih.Svc.PostNewMatching(matchCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "multiple":
			multipleCard := cards.MultipleChoice{}
			err = json.Unmarshal(data, &multipleCard)
			if err != nil {
				http.Error(w, "unable to decode database", http.StatusBadRequest)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err := validate.Struct(multipleCard)
			if err != nil {
				http.Error(w,"all fields require data", http.StatusBadRequest)
			}

			err = ih.Svc.PostNewMultiple(multipleCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "info":
			card := cards.Info{}
			err = json.Unmarshal(data, &card)
			if err != nil {
				http.Error(w,"unable to decode database", http.StatusUnprocessableEntity)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err := validate.Struct(card)
			if err != nil {
				http.Error(w,"all fields require data", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewInfo(card)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "qanda":
			qandaCard := cards.QNA{}
			err = json.Unmarshal(data, &qandaCard)
			if err != nil {
				http.Error(w, "unable to decode database", http.StatusBadRequest)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err := validate.Struct(qandaCard)
			if err != nil {
				http.Error(w,"all fields require data", http.StatusBadRequest)
				
			}

			err = ih.Svc.PostNewQNA(qandaCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "torf":
			torfCard := cards.TrueOrFalse{}
			err = json.Unmarshal(data, &torfCard)
			if err != nil {
				http.Error(w,"unable to decode database", http.StatusUnprocessableEntity)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err := validate.Struct(torfCard)
			if err != nil {
				http.Error(w,"all fields require data", http.StatusBadRequest)
			}

			err = ih.Svc.PostNewTORF(torfCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		default:
			http.Error(w, "invaild type", http.StatusBadRequest)
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (ih InfoHandler) GetFlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	myDb, err := ih.Svc.GetAllFlashcards()
	if err != nil {
		http.Error(w,"unable to decode database", http.StatusUnprocessableEntity)
	}

	db, err := json.MarshalIndent(myDb, "", " ")
	if err != nil {
		http.Error(w, "unable to encode database", http.StatusBadRequest)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(db)
}

func (ih InfoHandler) GetByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["type"]

	getType, err := ih.Svc.GetByType(id)
	if err != nil {
		switch err.Error() {
		case "type not found":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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

func (ih InfoHandler) DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := ih.Svc.DeletebyId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (ih InfoHandler) UpdateByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"unable to access database" , http.StatusUnprocessableEntity)
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusUnprocessableEntity)
	}

	err = ih.Svc.UpdatebyId(id, CardType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
		
	}

}
