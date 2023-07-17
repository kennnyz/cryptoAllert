package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6262902974:AAEAP9ZrTpVlKJtt9iZbzq9ArW2SPRN4R_w")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// handle the commands
		if update.Message.IsCommand() {
			handleCommand(update, bot)
			continue
		}
		if update.Message.Text == "/start" {
			handleStart(update, bot)
		} else if update.Message.Text == "Add Transfer" {
			handleAddTransfer(update, bot)
		} else if update.Message.Text == "Sell" {
			handleSell(update, bot)
		} else if update.Message.Text == "Buy" {
			handleBuy(update, bot)
		}
	}

}

// handleCommand handles bot commands

func handleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = "Hello! I am your bot. Use /help to see available commands."
		handleStart(update, bot)
	case "help":
		msg.Text = "Available commands:\n/addtransfer - Add a new transfer"
	case "addtransfer":
		handleAddTransfer(update, bot)
	default:
		msg.Text = "Unknown command. Use /help to see available commands."
	}
	bot.Send(msg)
}

func handleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! What would you like to do?")
	msg.ReplyMarkup = getMainMenuKeyboard()
	bot.Send(msg)
}

func handleAddTransfer(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please choose 'Sell' or 'Buy'")
	msg.ReplyMarkup = getTransferMenuKeyboard()
	bot.Send(msg)
}

func handleSell(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You clicked 'Sell' button")
	bot.Send(msg)
	// we want to go back to main menu
	handleStart(update, bot)
}

func handleBuy(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You clicked 'Buy' button")
	bot.Send(msg)
	// we want to go back to main menu
	handleStart(update, bot)
}

func getMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.NewKeyboardButton("Add Transfer"),
				tgbotapi.NewKeyboardButton("Add Crypto"),
				tgbotapi.NewKeyboardButton("My Wallet"),
			},
		},
	}
}

func getTransferMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.NewKeyboardButton("Sell"),
				tgbotapi.NewKeyboardButton("Buy"),
			},
		},
	}
}
