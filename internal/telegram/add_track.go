package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/models"
	"log"
	"strconv"
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

		var percent float64
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Write the trigger % after which you want to receive notifications. (Min 5%)")
		b.bot.Send(msg)
		for percentUpdate := range b.updatesChan {
			if percentUpdate.Message == nil {
				continue
			}

			selectedPercent := percentUpdate.Message.Text
			if selectedPercent == "" {
				continue
			}

			p, err := strconv.ParseFloat(percentUpdate.Message.Text, 64)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide correct")
				b.bot.Send(msg)
				continue
			}
			percent = p
			break
		}
		// Если получено сообщение от пользователя, сохраняем пару в базу данных и завершаем функцию
		coin := models.UserCoin{
			Name:    selectedPair,
			UserID:  update.Message.Chat.ID,
			Percent: percent,
		}

		//
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
