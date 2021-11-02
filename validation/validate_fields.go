package validation

import (
	"errors"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
)

func ValidateMatchingFields(card cards.Matching) error {

	return nil
}
func ValidateMultipleChoiceFields(card cards.MultipleChoice) error {

	return nil
}
func ValidateOnlyInfoFields(card cards.Info) error {
	if len(card.Category) <= 1 {
		return errors.New("category is not valid")
	}

	if len(card.Information) <= 3 {
		return errors.New("category is not valid")
	}

	return nil
}
func ValidateQNAFields(card cards.QNA) error {
	if len(card.Category) <= 1 {
		return errors.New("category is not valid")
	}

	if len(card.Question) <= 7 {
		return errors.New("category is not valid")
	}

	if len(card.Answer) <= 7 {
		return errors.New("category is not valid")
	}

	return nil
}
func ValidateTorFFields(card cards.TrueOrFalse) error {
	if len(card.Category) <= 2 {
		return errors.New("category is not valid")
	}

	if len(card.Question) <= 7 {
		return errors.New("category is not valid")
	}

	if !card.Answer {
		return errors.New("answer should be of type bool")
	}

	return nil
}
