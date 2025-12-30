package server

import (
	"encoding/json"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "server": "coco-sml"})
}
