package api

import (
	"encoding/json"
	"net/http"
	"github.com/a-digi/coco-sml/src/server/route"
	"github.com/a-digi/coco-sml/src/db/binary"
)

// RegisterRoutes registers all API endpoints using the provided RouteBuilder
func RegisterRoutes(rb *route.RouteBuilder) {
	// Example: /v1/status endpoint
	rb.Handle("/v1/status", statusHandler)
}

// statusHandler handles GET /v1/status requests
func statusHandler(w http.ResponseWriter, r *http.Request, _ *binary.Server) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "server": "coco-sml"})
}
