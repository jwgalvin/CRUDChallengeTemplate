package handlers

import (
	"net/http"

	"example.com/crudapp/internal/store"
)

type ItemsHandler struct {
	store store.Store
}

func NewItemsHandler(st store.Store) *ItemsHandler {
	panic("not implemented")
}

func (h *ItemsHandler) HandleItems(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *ItemsHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *ItemsHandler) listItems(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *ItemsHandler) getItem(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented")
}

func (h *ItemsHandler) createItem(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
