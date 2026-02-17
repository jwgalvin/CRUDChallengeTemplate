package api

import (
	"net/http"

	"example.com/crudapp/internal/handlers"
	"example.com/crudapp/internal/store"
)

type Server struct {
    mux *http.ServeMux
}

func NewServer(st store.Store) *Server {
    mux := http.NewServeMux()
    itemsHandler := handlers.NewItemsHandler(st)

    mux.HandleFunc("/health", handlers.Health)
    mux.HandleFunc("/items", itemsHandler.HandleItems)
    mux.HandleFunc("/items/", itemsHandler.HandleItem)

    return &Server{mux: mux}
}

func (s *Server) Handler() http.Handler {
    return s.mux
}
