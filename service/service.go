package service

import (
	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
)

type Repo interface {
	CreateNewInfo(card incomingdata.Info) error
	GetAllFlashcards() (repo.NewInfo, error) 
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) PostNewInfo(card incomingdata.Info) error {
	err := s.Repo.CreateNewInfo(card)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAllFlashcards() (repo.NewInfo, error){
	fc, err := s.Repo.GetAllFlashcards()
	if err != nil {
		return fc, err
	}
	return fc, nil
}
