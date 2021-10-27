package validation

import (
	"github.com/gpark1005/FlashCardsTeamOne/cards"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"errors"
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