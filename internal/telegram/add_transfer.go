package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleAddTransfer(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please choose 'Sell' or 'Buy'")
	msg.ReplyMarkup = getTransferMenuKeyboard()
	bot.Send(msg)
}
