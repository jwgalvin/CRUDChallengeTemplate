package handlers

import (
	"net/http"

	"example.com/crudapp/internal/store"
)

type ItemsHandler struct {
	store store.Store
}

func NewItemsHandler(st store.Store) *ItemsHandler {
	return &ItemsHandler{store: st}
}

func (h *ItemsHandler) HandleItems(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement route dispatch:
	// - GET /items → listItems()
	// - POST /items → createItem()
	// - Other methods → 405 Method Not Allowed
	panic("not implemented")
}

func (h *ItemsHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement route dispatch:
	// - Extract ID from URL path (e.g., /items/item-123 → id="item-123")
	// - Validate ID is not empty and has no nested paths
	// - GET /items/{id} → getItem()
	// - PUT /items/{id} → updateItem()
	// - DELETE /items/{id} → deleteItem()
	// - Other methods → 405 Method Not Allowed
	panic("not implemented")
}

func (h *ItemsHandler) listItems(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GET /items with filtering and pagination:
	// 1. Parse query parameters (limit, offset, name, tag) using validation.ParseListFilter()
	// 2. Call h.store.List() with parsed filter
	// 3. Return 200 OK with JSON: {items, limit, offset, count}
	// 4. Return 400 Bad Request if query params invalid
	// 5. Return 500 Internal Server Error on store error
	panic("not implemented")
}

func (h *ItemsHandler) getItem(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: Implement GET /items/{id}:
	// 1. Call h.store.Get(id)
	// 2. Return 200 OK with item JSON if found
	// 3. Return 404 Not Found if item doesn't exist
	// 4. Return 500 Internal Server Error on store error
	panic("not implemented")
}

func (h *ItemsHandler) createItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement POST /items:
	// 1. Parse JSON body into model.ItemInput using httpjson.DecodeJSON()
	// 2. Validate and normalize input using validation.NormalizeItemInput()
	// 3. Call h.store.Create() to save (store generates ID and timestamps)
	// 4. Return 201 Created with created item JSON
	// 5. Return 400 Bad Request if JSON invalid or validation fails
	// 6. Return 500 Internal Server Error on store error
	panic("not implemented")
}
