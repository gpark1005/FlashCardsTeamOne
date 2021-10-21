package main

import (
	"github.com/gpark1005/FlashCardsTeamOne/handler"
	"github.com/gpark1005/FlashCardsTeamOne/repo"
	"github.com/gpark1005/FlashCardsTeamOne/service"
)

func main() {
	fn := "/Users/austinnicholson/Desktop/GolandProjects/FlashcardTeam/FlashCardsTeamOne/flashcardsDb.json"

	// ext := filepath.Ext(fn)
	// if ext != ".json" {
	// 	log.Fatalln("File extension invaild")
	// }

	r := repo.NewRepo(fn)

	svc := service.NewService(r)

	hdlr := handler.NewInfoHandler(svc)

	router := handler.ConfigureRouter(hdlr)

	handler.NewServer(router)
}
