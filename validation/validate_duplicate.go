package validation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
)

type dbVal struct {
	Flashcards []map[string]interface{}
}

/* This function compares the questions map from the incoming data to
what is stored in the db to see if the card is already saved */
func ValidateDuplicateMatching(card cards.Matching) error {
	filename := "flashcardsDb.json"
	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("unable to read file")
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return errors.New("unable to decode database")
	}

	/* This loop with the help of the if ok statment, allows us to access the questions map stored on
	each matching card saved in the db. Since we are comparing two maps, reflect.DeepEqual will return
	true if they match or false if they do not. The bool is stored in value eq. */
	for _, val := range Db.Flashcards {
		if cQ, ok := val["questions"]; ok {
			eq := reflect.DeepEqual(card.Questions, cQ)
			if eq {
				return errors.New("this card already exists")
			}
		}
	}

	for _, val := range Db.Flashcards {
		if cQ, ok := val["options"]; ok {
			eq := reflect.DeepEqual(card.Answers, cQ)
			if eq {
				return errors.New("this card already exists")
			}
		}
	}

	return nil
}

/* This function compares the question from the incoming data to
what is stored in the db to see if the card is already saved */
func ValidateDuplicateMultipleChoice(card cards.MultipleChoice) error {
	filename := "flashcardsDb.json"
	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("error reading file")
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return errors.New("error decoding")
	}

	/* This loop with the help of the if ok statment, allows us to access the question string stored on
	each multiple choice card saved in the db. Since we are comparing two strings, the only thing needed
	to compare the two values is an additional is statment. */
	for _, val := range Db.Flashcards {
		if cQ, ok := val["question"]; ok {
			if cQ == card.Question {
				return errors.New("this card already exists")
			}
		}
	}
	return nil
}

func ValidateDuplicateOnlyInfo(card cards.Info) error {
	filename := "flashcardsDb.json"
	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("error reading file")
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return errors.New("error decoding")
	}

	for _, val := range Db.Flashcards {
		if cInfo, ok := val["information"]; ok {
			if cInfo == card.Information {
				return errors.New("this card already exists")
			}
		}
	}
	return nil
}

func ValidateDuplicateQNA(card cards.QNA) error {
	filename := "flashcardsDb.json"
	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("error reading file")
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return errors.New("error decoding")
	}

	for _, val := range Db.Flashcards {
		if cQ, ok := val["question"]; ok {
			if cQ == card.Question {
				return errors.New("this card already exists")
			}
		}
	}
	return nil
}

func ValidateDuplicateTorF(card cards.TrueOrFalse) error {
	filename := "flashcardsDb.json"
	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("error reading file")
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return errors.New("error decoding")
	}

	for _, val := range Db.Flashcards {
		if cQ, ok := val["question"]; ok {
			if cQ == card.Question {
				return errors.New("this card already exists")
			}
		}
	}
	return nil
}
