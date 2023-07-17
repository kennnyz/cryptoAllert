package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleSell(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You clicked 'Sell' button")
	bot.Send(msg)
	// we want to go back to main menu
	handleStart(update, bot)
}
