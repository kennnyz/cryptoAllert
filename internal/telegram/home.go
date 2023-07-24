package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) HandleCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = "Hello! I am your bot. Use /help to see available commands."
		handleStart(update, b.bot)
	case "help":
		msg.Text = "Available commands:\n/newtrack - Add a new coin to tracking\n/mytracks - Show your tracking coins\n/price - Show price of your tracking coins"
	case "newtrack":
		b.handleAddCoinToTrack(update)
	case "mytracks":
		handleMyTracks(update, b.bot)
	case "price":
		// TODO handle price of user coins
	default:
		msg.Text = "Unknown command. Use /help to see available commands."
	}
	b.bot.Send(msg)
}
