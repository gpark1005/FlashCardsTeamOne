package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"github.com/gpark1005/FlashCardsTeamOne/cards"
)

type Db struct {
	Flashcards []interface{}
}

type DbType struct {
	Flashcards []map[string]interface{}
}

type Repo struct {
	Filename string
}

func NewRepo(fn string) Repo {
	return Repo{
		Filename: fn,
	}
}

func (r Repo) CreateNewInfo(card cards.Info) error {
	newCardInfo := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newCardInfo)
	if err != nil {
		return err
	}

	newCardInfo.Flashcards = append(newCardInfo.Flashcards, card)

	input, err := json.MarshalIndent(newCardInfo, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) CreateNewMatching(card cards.Matching) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return err
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

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

func (r Repo) CreateNewMultiple(card cards.MultipleChoice) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return err
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

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

func (r Repo) CreateNewQNA(card cards.QNA) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return err
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

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

func (r Repo) CreateNewTORF(card cards.TrueOrFalse) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return err
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

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

func (r Repo) GetAllFlashcards() (Db, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return Db{}, err
	}

	flashcards := Db{}
	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, err
	}
	return flashcards, nil
}

func (r Repo) GetByType(input string) (DbType, error) {
	
	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return flashcards, err
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, err
	}


	for _, val := range flashcards.Flashcards {
		if cType, ok := val["type"]; ok {
			if cType == input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}

		}
	}

	return newDb, errors.New("type not found")
}