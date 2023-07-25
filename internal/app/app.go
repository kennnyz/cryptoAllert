package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	postgres2 "github.com/kennnyz/cryptoAllert/internal/storage/postgres"
	"github.com/kennnyz/cryptoAllert/internal/telegram"
	"github.com/kennnyz/cryptoAllert/pkg/postgres"
	"log"
)

func Run() {
	tBot, err := tgbotapi.NewBotAPI("6262902974:AAEAP9ZrTpVlKJtt9iZbzq9ArW2SPRN4R_w")
	if err != nil {
		log.Println(err)
		return
	}
	db, err := postgres.NewClient("user=postgres password=password dbname=telegram host=localhost port=5432 sslmode=disable")
	repo := postgres2.NewTelegramDB(db)

	bot := telegram.NewBot(tBot, repo)

	bot.Start()
}
