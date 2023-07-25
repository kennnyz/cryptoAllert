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

func (t *TelegramDB) AddUser(chatId int64) error {
	_, err := t.db.Exec("INSERT INTO users (id) VALUES ($1)", chatId)
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

func (t *TelegramDB) GetWallet(userID int64) ([]models.UserCoin, error) {
	rows, err := t.db.Query("SELECT name, user_id FROM user_coins WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userCoins []models.UserCoin

	for rows.Next() {
		var name string
		var userID int64

		err := rows.Scan(&name, &userID)
		if err != nil {
			return nil, err
		}

		userCoin := models.UserCoin{
			Name:   name,
			UserID: userID,
		}

		userCoins = append(userCoins, userCoin)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userCoins, nil
}
