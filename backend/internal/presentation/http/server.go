package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	return &Server{
		router: r,
	}
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func (s *Server) Start(port string) {
	log.Printf("Server is running on port %s", port)

	if err := http.ListenAndServe(":"+port, s.router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
