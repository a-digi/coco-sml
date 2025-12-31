package routes

import (
	"github.com/a-digi/coco-sml/src/api/domain/db"
	"github.com/a-digi/coco-sml/src/api/domain/status"
)

// RegisterRoutes registers all API endpoints using the provided RouteBuilder
func RegisterRoutes(rb *RouteBuilder) {
	// Example: /v1/status endpoint
	rb.Handle("/", status.StatusHandler)
	// Database: List tables endpoint
	rb.Handle("/v1/db/tables", db.DatabaseListTablesHandler)
}