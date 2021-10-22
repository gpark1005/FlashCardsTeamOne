package repo

import (
	"encoding/json"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)

// type NewInfo struct {
// 	Flashcard []incomingdata.Info
// }

type Repo struct {
	Filename string
}

func NewRepo(fn string) Repo {
	return Repo{
		Filename: fn,
	}
}

func (r Repo) CreateNewInfo(card incomingdata.Info) error {
	newcard := incomingdata.NewInfo{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return err
	}

	if card.Id == "" {
		card.Id = uuid.New().String()
	}

	newcard.Flashcard = append(newcard.Flashcard, card)

	input, err := json.MarshalIndent(newcard, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) GetAllFlashcards() (incomingdata.NewInfo, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return incomingdata.NewInfo{}, err
	}

	flashcards := incomingdata.NewInfo{}
	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, err
	}
	return flashcards, nil
}
