package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		saveSelectedPairToDatabase(selectedPair)
		mainMenuMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your pair has been saved. What would you like to do next?")
		mainMenuMsg.ReplyMarkup = getMainMenuKeyboard()
		b.bot.Send(mainMenuMsg)
		return
	}
}

func saveSelectedPairToDatabase(pair string) {
	// Здесь реализуйте логику для сохранения выбранной пары в базе данных
	// Например, выполните INSERT-запрос в базу данных, чтобы добавить выбранную пару в таблицу с отслеживаемыми парами
	// ...
	fmt.Printf("Selected pair to add: %s\n", pair)
}
