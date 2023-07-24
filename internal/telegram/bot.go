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
	bot         *tgbotapi.BotAPI
	updatesChan tgbotapi.UpdatesChannel
	repository  Repository
}

func NewBot(bot *tgbotapi.BotAPI, repository Repository) *Bot {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &Bot{
		bot:         bot,
		updatesChan: updates,
		repository:  repository,
	}
}

func (b *Bot) Start() {
	for update := range b.updatesChan {
		if update.Message == nil {
			continue
		}
		// handle commands
		if update.Message.IsCommand() {
			b.HandleCommand(update)
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
