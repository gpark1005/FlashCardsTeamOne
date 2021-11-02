package main

import (
	"github.com/gpark1005/FlashCardsTeamOne/handler"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
	"github.com/gpark1005/FlashCardsTeamOne/service"
)

func main() {
	fn := "flashcardsDb.json"

	r := repo.NewRepo(fn)

	svc := service.NewService(r)

	hdlr := handler.NewCardHandler(svc)

	router := handler.ConfigureRouter(hdlr)

	handler.NewServer(router)
}
