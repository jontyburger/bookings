package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jontyburger/bookings/pkg/config"
	"github.com/jontyburger/bookings/pkg/handlers"
)

// Routes function
func routes(app *config.AppConfig) http.Handler {

	// create a new handler with the pat.new()
	//mux := pat.New()

	mux := mux.NewRouter()

	// Home handler
	mux.HandleFunc("/", handlers.Repo.Home).Methods("GET")
	// About handler
	mux.HandleFunc("/about", handlers.Repo.About).Methods(("GET"))

	// return mux
	return mux

}
