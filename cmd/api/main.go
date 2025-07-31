package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/registaalfe/authuser-api-golang/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
  ____   ___   ____ ___ 
 / ___| / _ \ / ___|_ _|
| |  _ | | | | |    | | 
| |_| || |_| | |___ | | 
 \____(_)___/ \____|___|
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
