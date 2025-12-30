package binary

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// DataType represents the supported field types for schema validation.
type DataType int

const (
	StringType DataType = iota
	IntType
	FloatType
	BoolType
	DateType
	DateTimeType
	BinaryType
	UUIDType
	JSONType
)

func (dt DataType) String() string {
	switch dt {
	case StringType:
		return "string"
	case IntType:
		return "int"
	case FloatType:
		return "float"
	case BoolType:
		return "bool"
	case DateType:
		return "date"
	case DateTimeType:
		return "datetime"
	case BinaryType:
		return "binary"
	case UUIDType:
		return "uuid"
	case JSONType:
		return "json"
	default:
		return "unknown"
	}
}

// MarshalJSON serializes DataType as a string for JSON.
func (dt DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

// UnmarshalJSON deserializes DataType from a string in JSON.
func (dt *DataType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "string":
		*dt = StringType
	case "int":
		*dt = IntType
	case "float":
		*dt = FloatType
	case "bool":
		*dt = BoolType
	case "date":
		*dt = DateType
	case "datetime":
		*dt = DateTimeType
	case "binary":
		*dt = BinaryType
	case "uuid":
		*dt = UUIDType
	case "json":
		*dt = JSONType
	default:
		*dt = StringType // fallback or handle error
	}
	return nil
}

// Field represents a schema field for a table, including type and validation metadata.
type Field struct {
	Name      string   `json:"name"`      // Field name (unique within the table)
	DataType  DataType `json:"dataType"` // Data type (enum: StringType, IntType, FloatType, BoolType)
	MinLength int      `json:"minLength"`// Minimum length for string fields (0 if not applicable)
	MaxLength int      `json:"maxLength"`// Maximum length for string fields (0 if not applicable)
	Required  bool     `json:"required"` // Whether this field is required (must be present in every entry)
}

// Table represents metadata and schema information for a logical table in the search engine.
type Table struct {
	Name        string    `json:"name"`        // Unique name of the table
	Description string    `json:"description"` // Optional description of the table's purpose
	CreatedAt   time.Time `json:"createdAt"`   // Timestamp when the table was created
	UpdatedAt   time.Time `json:"updatedAt"`   // Timestamp when the table was last updated
	Fields      []Field   `json:"fields"`      // List of schema fields with type and validation metadata
	EntryCount  int       `json:"entryCount"`  // Number of entries currently in the table
}

var ErrTableExists = errors.New("table already exists")
var ErrInvalidTableName = errors.New("invalid table name")
var ErrMetaFileCorrupt = errors.New("metadata file corrupt")

// TablesMeta represents the central metadata for all tables in the database.
type TablesMeta struct {
	Tables []Table
}

// metaFileName is the name of the central metadata file.
const metaFileName = "tables.meta"

// CreateTable creates a new table by adding its metadata to the central binary metadata file in the given data directory.
// Returns an error if the table name is invalid, already exists, or writing fails.
func CreateTable(dataDir string, table Table) error {
	validName := regexp.MustCompile(`^[A-Za-z_]+$`)
	if !validName.MatchString(table.Name) {
		return ErrInvalidTableName
	}
	metaPath := filepath.Join(dataDir, metaFileName)

	var meta TablesMeta
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
