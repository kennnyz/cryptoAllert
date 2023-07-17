package postgres

import (
	"database/sql"
	"github.com/kennnyz/cryptoAllert/internal/models"
)

type TelegramDB struct {
	db *sql.DB
}

func NewTelegramDB(db *sql.DB) *TelegramDB {
	return &TelegramDB{db: db}
}

func (t *TelegramDB) AddUser(chatId int) error {
	_, err := t.db.Exec("INSERT INTO users (id) VALUES ($1)", chatId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TelegramDB) AddTransfer(transfer models.Transfer) error {
	query := `INSERT INTO transfers (type, amount, price, user_id) VALUES ($1, $2, $3, $4)`
	_, err := t.db.Exec(query, transfer.ActionType, transfer.Amount, transfer.Price, transfer.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TelegramDB) AddCoin(coin models.UserCoin) error {
	query := `INSERT INTO user_coins (name, user_id) VALUES ($1, $2)`
	_, err := t.db.Exec(query, coin.Name, coin.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TelegramDB) GetWallet(userID int) ([]models.UserCoin, error) {
	rows, err := t.db.Query("SELECT name, user_id, amount, usd_amount FROM user_coins WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userCoins []models.UserCoin

	for rows.Next() {
		var name string
		var userID int
		var amount float64
		var usdAmount float64

		err := rows.Scan(&name, &userID, &amount, &usdAmount)
		if err != nil {
			return nil, err
		}

		userCoin := models.UserCoin{
			Name:      name,
			UserID:    userID,
			Amount:    amount,
			USDAmount: usdAmount,
		}

		userCoins = append(userCoins, userCoin)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userCoins, nil
}
