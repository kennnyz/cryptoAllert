package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! What would you like to do?")
	msg.ReplyMarkup = getMainMenuKeyboard()
	bot.Send(msg)
}
