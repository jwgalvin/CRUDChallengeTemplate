package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/crudapp/internal/api"
	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store/sqlite"
)

func newHandler(t *testing.T) http.Handler {
	t.Helper()
	st := sqlite.New(":memory:")
	return api.NewServer(st).Handler()
}

func postItem(t *testing.T, handler http.Handler, name string, tags []string) model.Item {
	t.Helper()
	payload := map[string]interface{}{"name": name, "tags": tags}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusCreated {
		t.Fatalf("postItem: expected 201, got %d: %s", res.Code, res.Body.String())
	}
	var item model.Item
	if err := json.NewDecoder(res.Body).Decode(&item); err != nil {
		t.Fatalf("postItem: decode: %v", err)
	}
	return item
}


func TestCreateItem(t *testing.T) {
	handler := newHandler(t)
	item := postItem(t, handler, "alpha", []string{"go"})
	if item.ID == "" {
		t.Error("expected non-empty ID")
	}
	if item.Name != "alpha" {
		t.Errorf("expected name alpha, got %s", item.Name)
	}
}

func TestCreateItemNoTags(t *testing.T) {
	handler := newHandler(t)
	item := postItem(t, handler, "no-tags", nil)
	if item.Tags == nil {
		t.Error("expected empty slice, got nil")
	}
}

func TestCreateItemEmptyName(t *testing.T) {
	handler := newHandler(t)
	body, _ := json.Marshal(map[string]interface{}{"name": "  "})
	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.Code)
	}
}

func TestCreateItemInvalidBody(t *testing.T) {
	handler := newHandler(t)
	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader([]byte("not json")))
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.Code)
	}
}

func TestGetItem(t *testing.T) {
	handler := newHandler(t)
	created := postItem(t, handler, "beta", []string{"go"})

	req := httptest.NewRequest(http.MethodGet, "/items/"+created.ID, nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.Code)
	}
	var got model.Item
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected id %s, got %s", created.ID, got.ID)
	}
}

func TestGetItemNotFound(t *testing.T) {
	handler := newHandler(t)
	req := httptest.NewRequest(http.MethodGet, "/items/item-doesnotexist", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", res.Code)
	}
}

func TestListItems(t *testing.T) {
	handler := newHandler(t)
	postItem(t, handler, "one", nil)
	postItem(t, handler, "two", nil)

	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.Code)
	}
	var items []model.Item
	if err := json.NewDecoder(res.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestListItemsByName(t *testing.T) {
	handler := newHandler(t)
	postItem(t, handler, "gopher", nil)
	postItem(t, handler, "rustacean", nil)

	req := httptest.NewRequest(http.MethodGet, "/items?name=gopher", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	var items []model.Item
	json.NewDecoder(res.Body).Decode(&items)
	if len(items) != 1 || items[0].Name != "gopher" {
		t.Errorf("expected 1 result named gopher, got %d", len(items))
	}
}

func TestListItemsByTag(t *testing.T) {
	handler := newHandler(t)
	postItem(t, handler, "tagged", []string{"golang"})
	postItem(t, handler, "untagged", nil)

	req := httptest.NewRequest(http.MethodGet, "/items?tag=golang", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	var items []model.Item
	json.NewDecoder(res.Body).Decode(&items)
	if len(items) != 1 {
		t.Errorf("expected 1 result with tag golang, got %d", len(items))
	}
}

func TestListItemsPagination(t *testing.T) {
	handler := newHandler(t)
	for i := range 5 {
		postItem(t, handler, fmt.Sprintf("item-%d", i), nil)
	}

	req := httptest.NewRequest(http.MethodGet, "/items?limit=2&offset=0", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	var items []model.Item
	json.NewDecoder(res.Body).Decode(&items)
	if len(items) != 2 {
		t.Errorf("expected 2 items with limit=2, got %d", len(items))
	}
}

func TestListItemsInvalidLimit(t *testing.T) {
	handler := newHandler(t)
	req := httptest.NewRequest(http.MethodGet, "/items?limit=abc", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.Code)
	}
}

func TestMethodNotAllowedItems(t *testing.T) {
	handler := newHandler(t)
	for _, method := range []string{http.MethodPut, http.MethodDelete, http.MethodPatch} {
		req := httptest.NewRequest(method, "/items", nil)
		res := httptest.NewRecorder()
		handler.ServeHTTP(res, req)
		if res.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s /items: expected 405, got %d", method, res.Code)
		}
	}
}

func TestMethodNotAllowedItem(t *testing.T) {
	handler := newHandler(t)
	created := postItem(t, handler, "gamma", nil)
	for _, method := range []string{http.MethodPut, http.MethodDelete, http.MethodPatch} {
		req := httptest.NewRequest(method, "/items/"+created.ID, nil)
		res := httptest.NewRecorder()
		handler.ServeHTTP(res, req)
		if res.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s /items/{id}: expected 405, got %d", method, res.Code)
		}
	}
}
