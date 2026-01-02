package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// DataType represents the supported field types for schema validation.
type DataType int

const (
	// StringType entfernt
	IntType DataType = iota
	FloatType
	BoolType
	DateType
	DateTimeType
	BinaryType
	UUIDType
	JSONType
	VarcharType // hinzugefügt
)

var typeStringMap = map[DataType]string{
	IntType:      "int",
	FloatType:    "float",
	BoolType:     "bool",
	DateType:     "date",
	DateTimeType: "datetime",
	BinaryType:   "binary",
	UUIDType:     "uuid",
	JSONType:     "json",
	VarcharType:  "varchar",
}

var typeMap = map[string]DataType{
	"int":      IntType,
	"float":    FloatType,
	"bool":     BoolType,
	"date":     DateType,
	"datetime": DateTimeType,
	"binary":   BinaryType,
	"uuid":     UUIDType,
	"json":     JSONType,
	"varchar":  VarcharType,
}

func (dt DataType) String() string {
	if s, ok := typeStringMap[dt]; ok {
		return s
	}
	return "unknown"
}

func (dt DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

func (dt *DataType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	dType, err := ParseDataType(s)
	if err != nil {
		*dt = IntType
		return nil
	}

	*dt = dType
	return nil
}

// ParseDataType wandelt einen SQL-Feldtyp-String in einen DataType um.
func ParseDataType(fieldType string) (DataType, error) {
	fieldType = strings.ToLower(fieldType)
	if dt, ok := typeMap[fieldType]; ok {
		return dt, nil
	}

	return 0, fmt.Errorf("unsupported field type: %s", fieldType)
}

// Field represents a schema field for a table, including type and validation metadata.
type Field struct {
	Name         string      `json:"name"`
	DataType     DataType    `json:"dataType"`
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