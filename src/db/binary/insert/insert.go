package insert

import (
	"fmt"
	"github.com/a-digi/coco-sml/src/db/binary/sql"
)

// InsertToTable inserts the parsed InsertSQL payload into the target table
func InsertToTable(insertSQL *sql.InsertSQL) error {
	if insertSQL == nil {
		return fmt.Errorf("insertSQL is nil")
	}
	if insertSQL.Table == "" {
		return fmt.Errorf("table name is required")
	}
	if len(insertSQL.Values) == 0 {
		return fmt.Errorf("payload is empty")
	}

	// Step 2: Map fields and values to a row (as a map)
	row := make(map[string]interface{})
	for i, field := range insertSQL.Fields {
		if i < len(insertSQL.Values) {
			row[field] = insertSQL.Values[i]
		} else {
			row[field] = nil
		}
	}

	// Step 3: Insert into internal DB (placeholder)
	if err := insertRowToTable(insertSQL.Table, row); err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	return nil
}

// insertRowToTable inserts the parsed data into the internal DB (stub)
func insertRowToTable(table string, data map[string]interface{}) error {
	// TODO: Implement real DB insert logic
	fmt.Printf("[DB] Inserted into %s: %v\n", table, data)
	return nil
}
