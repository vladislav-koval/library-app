package book_http

import (
	"errors"
	"net/http"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /book", s.handlers.HandleCreateBook)
	mux.HandleFunc("PATCH /book/{id}", s.handlers.HandleUpdateBook)
	mux.HandleFunc("GET /book/{id}", s.handlers.HandleGetBook)
	mux.HandleFunc("GET /books", s.handlers.HandleFilterBooks)
	mux.HandleFunc("DELETE /book/{id}", s.handlers.HandleDeleteBook)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
