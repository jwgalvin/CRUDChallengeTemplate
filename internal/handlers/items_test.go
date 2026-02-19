package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/crudapp/internal/api"
	"example.com/crudapp/internal/store/sqlite"
)

type itemResponse struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func TestItemsHandlerCreateAndGet(t *testing.T) {
	st := sqlite.New(":memory:")
	server := api.NewServer(st)
	handler := server.Handler()

	payload := map[string]interface{}{
		"name": "alpha",
		"tags": []string{"one"},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(body))
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", res.Code)
	}

	var created itemResponse
	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/items/"+created.ID, nil)
	getRes := httptest.NewRecorder()

	handler.ServeHTTP(getRes, getReq)

	if getRes.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", getRes.Code)
	}
}
