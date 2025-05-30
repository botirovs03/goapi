package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"testuser": {
		AuthToken: "testtoken",
		Username:  "testuser",
	},
	"testuser2": {
		AuthToken: "testtoken2",
		Username:  "testuser2",
	},
	"testuser3": {
		AuthToken: "testtoken3",
		Username:  "testuser3",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"testuser": {
		Coins:    100,
		Username: "testuser",
	},
	"testuser2": {
		Coins:    200,
		Username: "testuser2",
	},
	"testuser3": {
		Coins:    300,
		Username: "testuser3",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {

	time.Sleep(1 * time.Second) // Simulate a delay

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(1 * time.Second) // Simulate a delay

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
