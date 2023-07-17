package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) HandleCommand() {
	msg := tgbotapi.NewMessage(b.update.Message.Chat.ID, "")
	switch b.update.Message.Command() {
	case "start":
		msg.Text = "Hello! I am your bot. Use /help to see available commands."
		handleStart(b.update, b.bot)
	case "help":
		msg.Text = "Available commands:\n/addtransfer - Add a new transfer\n/addcoin - Add new coin\n/mywallet - Show your wallet"
	case "addtransfer":
		handleAddTransfer(b.update, b.bot)
	default:
		msg.Text = "Unknown command. Use /help to see available commands."
	}
	b.bot.Send(msg)
}
