package handlers // Signify the package name is on handlers package

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"                                     // Package for change query parameters to struct
	"github.com/registaalfe/authuser-api-golang.git/api"            // Package for API definitions
	"github.com/registaalfe/authuser-api-golang.git/internal/tools" // Package for connecting to the database
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}              // Take the data from CoinBalanceParams struct
	var decoder *schema.Decoder = schema.NewDecoder() // To convert query parameters to struct
	var err error                                     // To handle errors

	err = decoder.Decode(&params, r.URL.Query()) // Save the query parameters to params struct

	// Check if the username is empty
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Get the databaseInterface and save it to database variable
	var database *tools.DatabaseInterface // Declare a variable to hold the database interface with pointer type
	database, err = tools.NewDatabase()   // Call func NewDatabase to get the database interface
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Get the data coins by get the username from the database
	var tokenDetails *tools.CoinDetails                      // Declare a variable as a pointer to CoinDetails struct
	tokenDetails = (*database).GetUserCoins(params.Username) // Get data from database
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json") // Set the data type in the response header to JSON
	err = json.NewEncoder(w).Encode(response)          // Changes struct Go to JSON format
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
