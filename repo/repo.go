package repo

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)

type NewInfo struct {
	Flashcard []incomingdata.Info
}

type Repo struct {
	Filename string
}

func NewRepo(fn string) Repo {
	return Repo{
		Filename: fn,
	}
}

func (r Repo) CreateNewInfo(Card incomingdata.Info) error {
	NI := NewInfo{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &NI)
	if err != nil {
		return err
	}

	NI.Flashcard = append(NI.Flashcard, Card)

	input, err := json.MarshalIndent(NI, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return err
	}
	return nil
}