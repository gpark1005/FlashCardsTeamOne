package validation

import (
	"errors"
	"strconv"

	"github.com/gpark1005/FlashCardsTeamOne/cards"
)

func ValidateMatchingFields(card cards.Matching) error {
	//checks to see if fields are valid and not blank
	if len(card.Category) <= 2 {
		return errors.New("category is not valid, must be more than 2 characters")
	}

	if len(card.Questions) == 0 {
		return errors.New("questions can not be left blank")
	}

	for _, v := range card.Questions {
		if len(v) <= 3 {
			return errors.New("question is not valid, must be more than 3 characters")
		}
	}

	if len(card.Options) == 0 {
		return errors.New("answers can not be left blank")
	}

	for _, v := range card.Options {
		if len(v) <= 3 {
			return errors.New("options is not valid, must be more than 3 characters")
		}
	}

	// check that will make sure question keys are a letter
	questionCount := 0

	for key := range card.Questions {
		_, err := strconv.Atoi(key)
		if err != nil {
			questionCount++
		}

		if questionCount != len(card.Questions) {
			return errors.New("each question key must be a letter")
		}
	}

	// check that will make sure answer keys are a letter
	aCount := 0

	for key := range card.Answers {
		_, err := strconv.Atoi(key)
		if err != nil {
			aCount++
		}

		if aCount != len(card.Answers) {
			return errors.New("each answer key must be a letter")
		}
	}

	// check that will make sure that the option keys are numbers that incroment up from 1
	optionCount := len(card.Options)
	okCount := 0
	for {
		if optionCount == 0 {
			break
		}

		check := card.Options
		countString := strconv.Itoa(optionCount)

		if _, ok := check[countString]; ok {
			if ok {
				okCount++
			}
		}
		optionCount--
	}

	if okCount != len(card.Options) {
		return errors.New("option keys must be a number that incroments up from 1. ex: 1, 2, 3, etc")
	}

	// check that will make sure answers are not left blank
	if len(card.Answers) == 0 {
		return errors.New("answers can not be left blank")
	}

	// check that will make sure answers are a number
	for _, v := range card.Answers {
		_, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("each answer needs to be a number")
		}
	}

	// check that will make sure answers are a option
	count := 0

	for _, v := range card.Answers {
		check := card.Options

		if _, ok := check[v]; ok {
			if ok {
				count++
			}
		}
	}

	if count != len(card.Answers) {
		return errors.New("answers must be a number that is in options")
	}

	// check that will make sure answer keys are a question key
	answerCount := 0

	for key := range card.Answers {
		check := card.Questions

		if _, ok := check[key]; ok {
			if ok {
				answerCount++
			}
		}
	}

	if answerCount != len(card.Answers) {
		return errors.New("answer keys must be a question key, case sensitive")
	}

	//checks to see if field names are valid
	if card.Questions == nil {
		return errors.New("field name must be questions")
	}

	if card.Options == nil {
		return errors.New("field name must be options")
	}

	if card.Answers == nil {
		return errors.New("field name must be answers")
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
