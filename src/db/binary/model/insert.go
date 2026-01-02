package model

import "time"

// InsertEvent represents a single insert event with metadata
// Payload must be a raw SQL string (e.g., an INSERT statement or SQL row data)
type InsertEvent struct {
	Table     string    // Table name
	Payload   string    // Raw SQL string for the insert operation
	Timestamp time.Time // Event timestamp
}
