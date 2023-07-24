package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleAddCoinToTrack(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide choose pare:")
	if update.Message == nil || update.Message.Text == "" {
		return
	}

	msg.ReplyMarkup = getAvailableCoinsKeyboard()
	selectedPair := update.Message.Text
	saveSelectedPairToDatabase(selectedPair)
	bot.Send(msg)
}

func saveSelectedPairToDatabase(pair string) {
	// Здесь реализуйте логику для сохранения выбранной пары в базе данных
	// Например, выполните INSERT-запрос в базу данных, чтобы добавить выбранную пару в таблицу с отслеживаемыми парами
	// ...
	fmt.Printf("Selected pair to add: %s\n", pair)
}
