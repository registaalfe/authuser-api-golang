package main

import (
	"fmt"
	"net/http" // provide HTTP client and server implementations

	"github.com/go-chi/chi"                                            // router for HTTP request routing like /users/{id}
	"github.com/registaalfe/authuser-api-golang.git/internal/handlers" // saving the logic every endpoint like /coin-balance
	log "github.com/sirupsen/logrus"                                   // library logging in Go to show log error/info with neat format
)

func main() {
	log.SetReportCaller(true) // Set the log to show the file and line number where the log was called

	var r *chi.Mux = chi.NewRouter() // Create a new router instance using chi
	handlers.Handler(r)              // Call the Handler function to register all routes and their handlers

	fmt.Println("Starting web on port 8000")
	err := http.ListenAndServe("localhost:8000", r) // Start the HTTP server on localhost:8000 with the router r
	if err != nil {
		log.Error(err)
	}
}
