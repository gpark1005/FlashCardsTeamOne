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
	GetByCategory(input string) (repo.DbType, error)
	GetById(input string) (repo.DbType, error)
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
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusUnprocessableEntity)
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

			//Validating that all fields in structs are field
			validate := validator.New()

			err = validate.Struct(matchCard)
			if err != nil {
				http.Error(w, "all fields require data: id: must be blank, type: type should match type of card, category: must be existing categeory, questions: make sure each question is filled, options: make sure each option is filled, answers: each answer should be filled and correspond to questions/options", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewMatching(matchCard)
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

			//Validating that all fields in structs are field
			validate := validator.New()

			err = validate.Struct(multipleCard)
			if err != nil {
				http.Error(w, "all fields require data: id: must be blank, type: type should match type of card, category: must be existing categeory, question: make sure the question is filled, options: make sure each option is filled, answer: answer should be filled and correspond to question/options", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewMultiple(multipleCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "info":
			card := cards.Info{}
			err = json.Unmarshal(data, &card)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err = validate.Struct(card)
			if err != nil {
				http.Error(w, "all fields require data: id: must be blank, type: type should match type of card, category: must be existing categeory, information: info should be filled", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewInfo(card)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "qanda":
			qandaCard := cards.QNA{}
			err = json.Unmarshal(data, &qandaCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err = validate.Struct(qandaCard)
			if err != nil {
				http.Error(w, "all fields require data: id: must be blank, type: type should match type of card, category: must be existing categeory, question: make sure the question is filled, answer: answer should be filled and correspond to question", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewQNA(qandaCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "torf":
			torfCard := cards.TrueOrFalse{}
			err = json.Unmarshal(data, &torfCard)
			if err != nil {
				http.Error(w, "request body syntax is not valid", http.StatusBadRequest)
			}

			//Validating that all fields in structs are field
			validate := validator.New()

			err = validate.Struct(torfCard)
			if err != nil {
				http.Error(w, "all fields require data: id: must be blank, type: type should match type of card, category: must be existing categeory, question: make sure the question is filled, answer: answer should be filled and correspond to question", http.StatusBadRequest)
				return
			}

			err = ih.Svc.PostNewTORF(torfCard)
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

func (ih InfoHandler) GetFlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	myDb, err := ih.Svc.GetAllFlashcards()
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

func (ih InfoHandler) GetByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["type"]

	getType, err := ih.Svc.GetByType(id)
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
		http.Error(w, "unable to access database", http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(data, &CardType)
	if err != nil {
		http.Error(w, "unable to decode database", http.StatusUnprocessableEntity)
		return
	}

	err = ih.Svc.UpdatebyId(id, CardType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (ih InfoHandler) GetByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	getCategory, err := ih.Svc.GetByCategory(category)
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

func (ih InfoHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	getId, err := ih.Svc.GetById(id)
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
