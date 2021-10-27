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

func ValidateMatching(card cards.Matching, filename string) error {

	Db := dbVal{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return err
	}

	for _, val := range Db.Flashcards {
		if cQ, ok := val["questions"]; ok {
			eq := reflect.DeepEqual(card.Questions, cQ)
			if eq {
				return errors.New("this card already exists")
			}
		}

	}

	return nil
}
