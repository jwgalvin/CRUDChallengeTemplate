package handlers

import (
	"net/http"

	"example.com/crudapp/internal/httpjson"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	httpjson.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
