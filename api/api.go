package api

// Import necessary packages

import (
	"encoding/json" // handle JSON to Go
	"net/http"      // handle HTTP requests/responses
)

// Defined The User Balance Params
type CoinBalanceParams struct { // Write "type" before defining a struct, "struct" its how to groupings of data
	Username string // Set the Username field to string type
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code will be 200
	Code int // Set the Code field to int type

	// Account Balance
	Balance int64 // Set the Balance field to int64 type
}

// Error Struct
type Error struct {
	// Error Code
	Code int // Set the Code field to int type

	// Error Message
	Message string // Set the Message field to string type
}

// Write a function to write the response in JSON format
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json") // Set the data type in the response header to JSON
	w.WriteHeader(code)                                // Set the HTTP status code in the response header 200,400, etc.

	json.NewEncoder(w).Encode(resp) // Changes struct Go to JSON format
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) { // Handle request errors by user
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) { // Handel internal server errors
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
