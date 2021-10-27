package validation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
)



func ValidateMatching(card cards.Matching, filename string) error {

	Db := cards.Matching{}

	compare, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(compare, &Db)
	if err != nil {
		return err
	}

	eq := reflect.DeepEqual(card.Questions, Db.Questions)
	if eq {
		return errors.New("this card already exists")
	}
	return nil
}