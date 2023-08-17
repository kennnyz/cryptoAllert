package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (b *Bot) handleMyTracks(update tgbotapi.Update) {
	// want to get all pairs that user have
	coins, err := b.repository.GetWallet(update.Message.Chat.ID)
	if err != nil {
		log.Println("Error getting user coins ", err)
	}

	s := strings.Builder{}
	for i, coin := range coins {
		s.WriteString(fmt.Sprintf("%d. %s\n", i+1, coin.Name))
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, s.String())

	b.bot.Send(msg)
}
