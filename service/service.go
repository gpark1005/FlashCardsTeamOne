package service

import (
	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)
type Repo interface {
	CreateNewInfo(Card incomingdata.Info) error 
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) PostNewInfo(Card incomingdata.Info) error {
	err := s.Repo.CreateNewInfo(Card)
	if err != nil {
		return err
	}
	return nil
}