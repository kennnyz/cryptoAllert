package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/models"
)

type Repository interface {
	AddUser(chatId int) error
	AddTransfer(transfer models.Transfer) error
	AddCoin(coin models.UserCoin) error
	GetWallet(userID int) ([]models.UserCoin, error)
}

type Bot struct {
	bot        *tgbotapi.BotAPI
	update     tgbotapi.Update
	repository Repository
}

func NewBot(bot *tgbotapi.BotAPI, update tgbotapi.Update, repository Repository) *Bot {
	return &Bot{
		bot:        bot,
		update:     update,
		repository: repository,
	}
}
