package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/registaalfe/authuser-api-golang.git/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// Public routes
	r.Get("/", homeHandler) // Register the root route with a handler
	r.Get("/hello", helloHandler)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Its Home Page")) // Write a simple response to the HTTP request
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World this is golang!")) // Write a simple response to the HTTP request
}
