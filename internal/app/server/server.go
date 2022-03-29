package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/KokoulinM/go-musthave-shortener-tpl/internal/app/handlers"
)

type Server struct {
	host string
}

func New(host string) *Server {
	return &Server{
		host: host,
	}
}

func (s *Server) Start() {
	handlers := handlers.New()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/", func(r chi.Router) {
		router.Get("/{id}", handlers.Get)
		router.Post("/", handlers.Save)
	})

	log.Fatal(http.ListenAndServe(s.host, router))
}