package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.NewKeyboardButton("Add Transfer"),
				tgbotapi.NewKeyboardButton("Add to track crypto"),
				tgbotapi.NewKeyboardButton("My tracking coins"),
			},
		},
	}
}

func getAvailablePairs() []string {
	// Вернуть список доступных пар, например ["BTC/USD", "ETH/USD", "XRP/USD", ...]
	// Здесь вы можете получить доступные пары из вашего источника данных (например, базы данных или API)
	return []string{"BTC/USD", "ETH/USD", "XRP/USD", "LTC/USD", "BCH/USD", "XLM/USD", "ADA/USD", "LINK/USD"}
}

func getAvailableCoinsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	availablePairs := getAvailablePairs()
	keyboard := make([][]tgbotapi.KeyboardButton, len(availablePairs))
	for i, pair := range availablePairs {
		btn := tgbotapi.NewKeyboardButton(pair)

		keyboard[i] = []tgbotapi.KeyboardButton{btn}
	}

	return tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}
