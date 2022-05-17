package service

import (
	"github.com/google/uuid"
	"github.com/gpark1005/FlashCardsTeamOne/cards"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
	"github.com/gpark1005/FlashCardsTeamOne/validation"
)

type Repo interface {
	CreateNewMatching(card cards.Matching) error
	CreateNewMultipleChoice(card cards.MultipleChoice) error
	CreateNewInfo(card cards.Info) error
	CreateNewQNA(card cards.QNA) error
	CreateNewTORF(card cards.TrueOrFalse) error
	GetAllFlashcards() (repo.Db, error)
	GetByType(input string) (repo.DbType, error)
	Delete(input string) error
	Update(input string, card map[string]interface{}) error
	GetByCategory(input string) (repo.DbType, error)
	GetById(input string) (repo.DbType, error)
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) PostNewMatching(card cards.Matching) error {
	card.Id = uuid.New().String()

	err := validation.ValidateMatchingFields(card)
	if err != nil {
		return err
	}

	err = validation.ValidateDuplicateMatching(card)
	if err != nil {
		return err
	}

	err = s.Repo.CreateNewMatching(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) PostNewMultipleChoice(card cards.MultipleChoice) error {
	card.Id = uuid.New().String()

	err := validation.ValidateMultipleChoiceFields(card)
	if err != nil {
		return err
	}

	err = validation.ValidateDuplicateMultipleChoice(card)
	if err != nil {
		return err
	}

	err = s.Repo.CreateNewMultipleChoice(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) PostNewInfo(card cards.Info) error {
	card.Id = uuid.New().String()

	err := validation.ValidateOnlyInfoFields(card)
	if err != nil {
		return err
	}

	err = validation.ValidateDuplicateOnlyInfo(card)
	if err != nil {
		return err
	}

	err = s.Repo.CreateNewInfo(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) PostNewQNA(card cards.QNA) error {
	card.Id = uuid.New().String()

	err := validation.ValidateQNAFields(card)
	if err != nil {
		return err
	}

	err = validation.ValidateDuplicateQNA(card)
	if err != nil {
		return err
	}

	err = s.Repo.CreateNewQNA(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) PostNewTORF(card cards.TrueOrFalse) error {
	card.Id = uuid.New().String()

	err := validation.ValidateTorFFields(card)
	if err != nil {
		return err
	}

	err = validation.ValidateDuplicateTorF(card)
	if err != nil {
		return err
	}

	err = s.Repo.CreateNewTORF(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAllFlashcards() (repo.Db, error) {
	fc, err := s.Repo.GetAllFlashcards()
	if err != nil {
		return fc, err
	}
	return fc, nil
}

func (s Service) GetByType(input string) (repo.DbType, error) {
	searchRequest, err := s.Repo.GetByType(input)
	if err != nil {
		return searchRequest, err
	}
	return searchRequest, nil
}

func (s Service) DeleteById(input string) error {
	err := s.Repo.Delete(input)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateById(input string, card map[string]interface{}) error {
	err := s.Repo.Update(input, card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetByCategory(input string) (repo.DbType, error) {
	searchRequest, err := s.Repo.GetByCategory(input)
	if err != nil {
		return searchRequest, err
	}
	return searchRequest, nil
}

func (s Service) GetById(input string) (repo.DbType, error) {
	searchRequest, err := s.Repo.GetById(input)
	if err != nil {
		return searchRequest, err
	}
	return searchRequest, nil
}
