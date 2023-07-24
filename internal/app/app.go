package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/telegram"
	"log"
)

func Run() {
	tBot, err := tgbotapi.NewBotAPI("6262902974:AAEAP9ZrTpVlKJtt9iZbzq9ArW2SPRN4R_w")
	if err != nil {
		log.Println(err)
		return
	}

	bot := telegram.NewBot(tBot, nil)

	bot.Start()
}
