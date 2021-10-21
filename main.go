package main

import (
	"fmt"
	"path/filepath"
	"log"
	"github.com/gpark1005/FlashCardsTeamOne/handler"
)



func main() {
	fmt.Println("Hello team!")

	fn := "FlashCardsTeamOne\flashcardsDb.json"

	ext := filepath.Ext(fn)
	if ext != ".json" {
		log.Fatalln("File extension invaild")
	}

	r := repo.NewRepo(fn)

	svc := service.NewService(r)

	hdlr := handler.NewInfoHandler(svc)

	router := handler.ConfigureRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())
}