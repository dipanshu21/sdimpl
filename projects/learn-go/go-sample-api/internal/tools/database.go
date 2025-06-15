package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username string
	Token    string
}

type CoinDetails struct {
	Username string
	Coins    int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (DatabaseInterface, error) {
	// Here you would typically connect to your database.
	// For simplicity, we return a mock implementation.
	db := &mockDB{}
	if err := db.SetupDatabase(); err != nil {
		log.Error("Failed to setup database:", err)
		return nil, err
	}
	return db, nil
}
