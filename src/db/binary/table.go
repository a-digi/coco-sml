package binary

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/a-digi/coco-sml/src/db/binary/model"
)

var ErrTableExists = errors.New("table already exists")
var ErrInvalidTableName = errors.New("invalid table name")
var ErrMetaFileCorrupt = errors.New("metadata file corrupt")

const metaFileName = "tables.meta"

// CreateTable creates a new table by adding its metadata to the central binary metadata file in the given data directory.
// Returns an error if the table name is invalid, already exists, or writing fails.
func CreateTable(dataDir string, table model.Table) error {
	validName := regexp.MustCompile(`^[A-Za-z_]+$`)
	if !validName.MatchString(table.Name) {
		return ErrInvalidTableName
	}
	metaPath := filepath.Join(dataDir, metaFileName)

	var meta model.TablesMeta
	if f, err := os.Open(metaPath); err == nil {
		defer f.Close()
		dec := gob.NewDecoder(f)
		if err := dec.Decode(&meta); err != nil {
			return ErrMetaFileCorrupt
		}
		for _, t := range meta.Tables {
			if t.Name == table.Name {
				return ErrTableExists
			}
		}
	}
	table.CreatedAt = time.Now()
	table.UpdatedAt = table.CreatedAt
	meta.Tables = append(meta.Tables, table)
	tmpPath := metaPath + ".tmp"
	file, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	if err := enc.Encode(meta); err != nil {
		return err
	}
	if err := os.Rename(tmpPath, metaPath); err != nil {
		return err
	}
	return nil
}

// TableMeta represents the metadata for a table (name and fields).
type TableMeta struct {
	Name   string
	Fields []struct {
		Name string
	}
}

// LoadTableMeta loads the metadata for all tables from tables.meta in the given dataDir.
func LoadTableMeta(dataDir string) ([]TableMeta, error) {
	metaPath := filepath.Join(dataDir, "tables.meta")
	metaFile, err := os.Open(metaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open metadata: %w", err)
	}
	defer metaFile.Close()
	var meta struct {
		Tables []TableMeta
	}
	if err := gob.NewDecoder(metaFile).Decode(&meta); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %w", err)
	}

	return meta.Tables, nil
}

// GetTableFields returns the field names for a given table name from the loaded metadata.
func GetTableFields(tables []TableMeta, tableName string) ([]string, error) {
	for _, t := range tables {
		if t.Name == tableName {
			var fields []string
			for _, f := range t.Fields {
				fields = append(fields, f.Name)
			}
			return fields, nil
		}
	}

	return nil, fmt.Errorf("table '%s' not found", tableName)
}

// ListTables returns all tables (with full metadata) in the given data directory.
func ListTables(dataDir string) ([]model.Table, error) {
	metaPath := filepath.Join(dataDir, metaFileName)
	metaFile, err := os.Open(metaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open metadata: %w", err)
	}
	defer metaFile.Close()

	var meta model.TablesMeta
	if err := gob.NewDecoder(metaFile).Decode(&meta); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %w", err)
	}
	return meta.Tables, nil
}

// TableMetaSliceToInterface converts a []TableMeta to a []interface{} for use in Result.
func TableMetaSliceToInterface(metas []TableMeta) []interface{} {
	results := make([]interface{}, len(metas))
	for i, m := range metas {
		results[i] = m
	}
	return results
}
