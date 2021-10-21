package service

import "github.com/gpark1005/FlashCardsTeamOne/incomingdata"

type Repo interface {

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
	
}