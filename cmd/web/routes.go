package main

import (
	"github/saifultechie/booking/pkg/config"
	"github/saifultechie/booking/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general", handlers.Repo.General)
	mux.Get("/major", handlers.Repo.Major)
	mux.Get("/search-availability", handlers.Repo.Search)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/search-availability-json", handlers.Repo.AvailabilityJson)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.ReservationPost)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("../../static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return mux
}
