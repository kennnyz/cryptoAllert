package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kennnyz/cryptoAllert/internal/models"
	"log"
)

type Repository interface {
	AddUser(chatId int64) error
	AddTransfer(transfer models.Transfer) error
	AddCoin(coin models.UserCoin) error
	GetWallet(userID int64) ([]models.UserCoin, error)
}

type Bot struct {
	bot        *tgbotapi.BotAPI
	update     tgbotapi.Update
	repository Repository
}

func NewBot(bot *tgbotapi.BotAPI, repository Repository) *Bot {
	return &Bot{
		bot:        bot,
		repository: repository,
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		// handle commands
		if update.Message.IsCommand() {
			b.HandleCommand(update.Message)
			continue
		}

		// adding new user
		if update.Message.NewChatMembers != nil {
			for _, newUser := range update.Message.NewChatMembers {
				chatID := newUser.ID
				err := b.repository.AddUser(chatID)
				if err != nil {
					log.Println(err)
					break
				}
			}
		}
	}
}
