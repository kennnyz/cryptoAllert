package models

type Transfer struct {
	UserID     int64   // chat id of the user
	ActionType string  // sell or buy
	Coin       string  // coin name
	Amount     float64 // amount of coins
	Price      float64 // price of the coin
}

type UserCoin struct {
	Name      string
	UserID    int64
	Amount    float64
	USDAmount float64
}
