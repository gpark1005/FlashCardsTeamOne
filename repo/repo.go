package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"

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

func (r Repo) CreateNewMatching(card cards.Matching) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return errors.New("unable to decode database")
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

	input, err := json.MarshalIndent(newcard, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil
}

func (r Repo) CreateNewMultipleChoice(card cards.MultipleChoice) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return errors.New("unable to decode database")
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

	input, err := json.MarshalIndent(newcard, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil
}

func (r Repo) CreateNewInfo(card cards.Info) error {
	newCardInfo := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(output, &newCardInfo)
	if err != nil {
		return errors.New("unable to decode database")
	}

	newCardInfo.Flashcards = append(newCardInfo.Flashcards, card)

	input, err := json.MarshalIndent(newCardInfo, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil
}

func (r Repo) CreateNewQNA(card cards.QNA) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return errors.New("unable to decode database")
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

	input, err := json.MarshalIndent(newcard, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil
}

func (r Repo) CreateNewTORF(card cards.TrueOrFalse) error {
	newcard := Db{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(output, &newcard)
	if err != nil {
		return errors.New("unable to decode database")
	}

	newcard.Flashcards = append(newcard.Flashcards, card)

	input, err := json.MarshalIndent(newcard, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil
}

func (r Repo) GetAllFlashcards() (Db, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return Db{}, errors.New("unable to read file")
	}

	flashcards := Db{}
	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, errors.New("unable to decode database")
	}
	return flashcards, nil
}

func (r Repo) GetByType(input string) (DbType, error) {

	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return flashcards, errors.New("unable to read file")
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, errors.New("unable to decode database")
	}

	for _, val := range flashcards.Flashcards {
		if cType, ok := val["type"]; ok {
			if cType == input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}

		}
	}

	mt := DbType{}

	eq := reflect.DeepEqual(mt.Flashcards, newDb.Flashcards)
	if eq {
		return newDb, errors.New("could not find that type in the data base")
	}

	return newDb, nil
}

func (r Repo) Delete(input string) error {

	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return errors.New("unable to decode database")
	}

	for _, val := range flashcards.Flashcards {
		if cId, ok := val["id"]; ok {
			if cId != input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}

		}
	}

	bytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, bytes, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}

	return nil
}

func (r Repo) Update(input string, card map[string]interface{}) error {

	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return errors.New("unable to decode database")
	}

	for _, val := range flashcards.Flashcards {
		if cId, ok := val["id"]; ok {
			if cId != input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}
			if cId == input {
				card["id"] = input
				newDb.Flashcards = append(newDb.Flashcards, card)
			}

		}
	}

	bytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return errors.New("unable to encode database")
	}

	err = ioutil.WriteFile(r.Filename, bytes, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}

	return nil
}

func (r Repo) GetByCategory(input string) (DbType, error) {

	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return flashcards, errors.New("unable to read file")
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, errors.New("unable to decode database")
	}

	for _, val := range flashcards.Flashcards {
		if cCategory, ok := val["category"]; ok {
			if cCategory == input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}

		}
	}
	mt := DbType{}

	eq := reflect.DeepEqual(mt.Flashcards, newDb.Flashcards)
	if eq {
		return newDb, errors.New("could not find that category in the data base")
	}

	return newDb, nil
}

func (r Repo) GetById(input string) (DbType, error) {

	flashcards := DbType{}
	newDb := DbType{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return flashcards, errors.New("unable to read file")
	}

	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return flashcards, errors.New("unable to decode database")
	}

	for _, val := range flashcards.Flashcards {
		if cId, ok := val["id"]; ok {
			if cId == input {
				newDb.Flashcards = append(newDb.Flashcards, val)
			}

		}
	}
	mt := DbType{}

	eq := reflect.DeepEqual(mt.Flashcards, newDb.Flashcards)
	if eq {
		return newDb, errors.New("could not find that id in the data base")
	}

	return newDb, nil
}
