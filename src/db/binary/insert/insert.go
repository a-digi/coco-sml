package insert

import (
	"fmt"
	"github.com/a-digi/coco-sml/src/db/binary/model"
)

// InsertToTable inserts the event payload into the target table
func InsertToTable(event model.InsertEvent) error {
	// Step 1: Validate payload
	if event.Table == "" {
		return fmt.Errorf("table name is required")
	}

	if event.Payload == "" {
		return fmt.Errorf("payload is empty")
	}

	// Step 2: Parse payload (placeholder)
	parsedData, err := parseInsertPayload(event.Payload)
	if err != nil {
		return fmt.Errorf("failed to parse payload: %w", err)
	}

	// Step 3: Insert into internal DB (placeholder)
	if err := insertRowToTable(event.Table, parsedData); err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	return nil
}

// parseInsertPayload parses the payload into a generic map (stub)
func parseInsertPayload(payload string) (map[string]interface{}, error) {
	// TODO: Implement real parsing logic
	// For now, just return a dummy map
	return map[string]interface{}{"raw": payload}, nil
}

// insertRowToTable inserts the parsed data into the internal DB (stub)
func insertRowToTable(table string, data map[string]interface{}) error {
	// TODO: Implement real DB insert logic
	fmt.Printf("[DB] Inserted into %s: %v\n", table, data)
	return nil
}
