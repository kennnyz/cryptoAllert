package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleBuy(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You clicked 'Buy' button")
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
	// we want to go back to main menu
	handleStart(update, bot)
}
