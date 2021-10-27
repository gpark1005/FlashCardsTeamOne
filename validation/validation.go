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
what is stored in the db to see if card is already saved */
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

	return nil
}
