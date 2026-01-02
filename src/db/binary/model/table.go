package model

import (
	"encoding/json"
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
	VarcharType // hinzugefügt
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
	case VarcharType:
		return "varchar"
	default:
		return "unknown"
	}
}

func (dt DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

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
	case "varchar":
		*dt = VarcharType
	default:
		*dt = StringType
	}
	return nil
}

// Field represents a schema field for a table, including type and validation metadata.
type Field struct {
	Name         string      `json:"name"`
	DataType     DataType    `json:"dataType"`
	MinLength    int         `json:"minLength"`
	MaxLength    int         `json:"maxLength"`
	Nullable     bool        `json:"nullable"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
}

// Table represents metadata and schema information for a logical table in the search engine.
type Table struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Fields      []Field   `json:"fields"`
	EntryCount  int       `json:"entryCount"`
}

// TablesMeta represents the central metadata for all tables in the database.
type TablesMeta struct {
	Tables []Table
}

// TableMeta represents the metadata for a table (name and fields).
type TableMeta struct {
	Name   string
	Fields []struct {
		Name string
	}
}