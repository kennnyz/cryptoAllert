package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/models"
)

func (b *Bot) handleAddCoin(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(b.update.Message.Chat.ID, "Please enter the coin name")

	// как получить сообщение от пользователя когда он находится внутри кнопки "Add coin"?
	coin := models.UserCoin{
		UserID: msg.ChatID,
		Name:   msg.Text,
	}
	b.repository.AddCoin(coin)
	b.bot.Send(msg)
}
