package repo

import (
	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)

type NewInfo struct {
	Flashcard incomingdata.Info
}

type Repo struct {
	Filename string
}

func NewRepo (fn string) Repo {
	return Repo {
		Filename: fn,
	}
}

