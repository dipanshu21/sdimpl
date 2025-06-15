package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]*LoginDetails{
	"alice": {
		Username: "alice",
		Token:    "alice-token",
	},
	"bob": {
		Username: "bob",
		Token:    "bob-token",
	},
	"charlie": {
		Username: "charlie",
		Token:    "charlie-token",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"alice": {
		Username: "alice",
		Coins:    1000,
	},
	"bob": {
		Username: "bob",
		Coins:    2000,
	},
	"charlie": {
		Username: "charlie",
		Coins:    3000,
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second) // Simulate a delay for the mock database call

	if details, exists := mockLoginDetails[username]; exists {
		return details
	}
	return nil
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second) // Simulate a delay for the mock database call

	if details, exists := mockCoinDetails[username]; exists {
		return details
	}
	return nil
}

func (db *mockDB) SetupDatabase() error {
	// In a real application, this would set up the database connection.
	// For the mock implementation, we just return nil to indicate success.
	return nil
}
