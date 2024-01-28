package cmd

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

)

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(middleware.Timeout(60))
		router.Use(middleware.Heartbeat("/ping"))
		router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		})
	return router
}

