package api

// Import necessary packages

// Defined The User Balance Params
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success Code will be 200
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}