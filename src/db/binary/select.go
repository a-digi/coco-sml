package sql

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ResultRow represents a single result row from a SELECT query.
type ResultRow map[string]interface{}

// ExecuteSelect executes a parsed SELECT SQLQuery and returns results as a slice of maps (field -> value).
// dataDir is the directory where tables.meta and <table>.data are stored.
func ExecuteSelect(query *SQLQuery, dataDir string) ([]ResultRow, error) {
	if query == nil {
		return nil, errors.New("nil query")
	}
	// Load table metadata using shared logic
	tables, err := LoadTableMeta(dataDir)
	if err != nil {
		return nil, err
	}
	tableFields, err := GetTableFields(tables, query.Table)
	if err != nil {
		return nil, err
	}

	// Open table data file (assume binary gob, e.g. <table>.data)
	dataPath := filepath.Join(dataDir, query.Table+".data")
	dataFile, err := os.Open(dataPath)

	if err != nil {
		return nil, fmt.Errorf("failed to open table data: %w", err)
	}

	defer dataFile.Close()
	dec := gob.NewDecoder(dataFile)
	var allRows []map[string]interface{}

	if err := dec.Decode(&allRows); err != nil {
		return nil, fmt.Errorf("failed to decode table data: %w", err)
	}

	// Filter rows by query.Conditions
	var filtered []ResultRow
	for _, row := range allRows {
		match := true
		for k, v := range query.Conditions {
			if fmt.Sprint(row[k]) != v {
				match = false
				break
			}
		}
		if match {
			filtered = append(filtered, row)
		}
	}

	// Select requested fields
	var results []ResultRow
	for _, row := range filtered {
		result := ResultRow{}
		if len(query.Fields) == 1 && query.Fields[0] == "*" {
			for _, f := range tableFields {
				result[f] = row[f]
			}
		} else {
			for _, f := range query.Fields {
				result[f] = row[f]
			}
		}
		results = append(results, result)
	}

	return results, nil
}
