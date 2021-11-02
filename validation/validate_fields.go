package validation

import (
	"errors"
	"fmt"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
)

func ValidateMatchingFields(card cards.Matching) error {
	var empty map[string]interface{}

	fmt.Printf("this is empty: %v", empty)

	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 characters")
	}

	return nil
}
func ValidateMultipleChoiceFields(card cards.MultipleChoice) error {
	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 characters")
	}

	if len(card.Question) <= 7 {
		return errors.New("question is not valid, must be more than 7 characters")
	}

	if len(card.Options.One) <= 2 {
		return errors.New("option one is not valid, must be more than 2 characters")
	}

	if len(card.Options.Two) <= 2 {
		return errors.New("option two is not valid, must be more than 2 characters")
	}

	if len(card.Options.Three) <= 2 {
		return errors.New("option three is not valid, must be more than 2 characters")
	}

	if len(card.Options.Four) <= 2 {
		return errors.New("option four is not valid, must be more than 2 characters")
	}

	if card.Answers >= 5 {
		return errors.New("answers is not valid, should be a number 1-4")
	}

	return nil
}
func ValidateOnlyInfoFields(card cards.Info) error {
	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 characters")
	}

	if len(card.Information) <= 7 {
		return errors.New("information is not valid, must be more than 7 characters")
	}

	return nil
}
func ValidateQNAFields(card cards.QNA) error {
	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 character long")
	}

	if len(card.Question) <= 7 {
		return errors.New("question is not valid, must be more than 7 characters")
	}

	if len(card.Answer) <= 7 {
		return errors.New("answer is not valid, must be more than 7 characters")
	}

	return nil
}
func ValidateTorFFields(card cards.TrueOrFalse) error {
	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 characters")
	}

	if len(card.Question) <= 7 {
		return errors.New("category is not valid, must be more than 7 characters")
	}

	if !card.Answer {
		return errors.New("answer should be of type bool")
	}

	return nil
}
