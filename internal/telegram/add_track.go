package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/models"
	"log"
)

func (b *Bot) handleAddCoinToTrack(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide choose pare:")
	msg.ReplyMarkup = getAvailableCoinsKeyboard()
	b.bot.Send(msg)

	for update := range b.updatesChan {
		// Проверяем, что получено сообщение от пользователя
		if update.Message == nil {
			continue
		}

		selectedPair := update.Message.Text
		if selectedPair == "" {
			continue
		}

		// Если получено сообщение от пользователя, сохраняем пару в базу данных и завершаем функцию
		coin := models.UserCoin{
			Name:   selectedPair,
			UserID: update.Message.Chat.ID,
		}

		err := b.repository.AddCoin(coin)
		if err != nil {
			log.Println("Err saving pair ", err)
		}

		mainMenuMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your pair has been saved. What would you like to do next?")
		mainMenuMsg.ReplyMarkup = getMainMenuKeyboard()
		b.bot.Send(mainMenuMsg)
		return
	}
}
