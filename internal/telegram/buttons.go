package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
