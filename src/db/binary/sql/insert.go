package sql

import (
	"fmt"
	"github.com/a-digi/coco-sml/src/db/binary/model"
)

// InsertSQL represents a parsed INSERT SQL statement.
type InsertSQL struct {
	Table  string
	Fields []string
	Values []interface{}
}

// ParseInsertSQL parses a basic INSERT SQL statement using the lexer and returns InsertSQL.
// Example: INSERT INTO users (id, name, age) VALUES (1, 'Alice', 30);
func ParseInsertSQL(query string, lexer func(string) []interface{}) (*InsertSQL, error) {
	tokens := lexer(query)
	pos := 0
	consume := func(expected string) bool {
		if pos < len(tokens) && tokens[pos].(map[string]interface{})["Type"] == expected {
			pos++
			return true
		}
		return false
	}
	// INSERT
	if !consume("INSERT") {
		return nil, fmt.Errorf("expected INSERT keyword")
	}
	// INTO
	if !consume("INTO") {
		return nil, fmt.Errorf("expected INTO keyword")
	}
	// table name
	if tokens[pos].(map[string]interface{})["Type"] != "IDENT" {
		return nil, fmt.Errorf("expected table name")
	}
	table := tokens[pos].(map[string]interface{})["Value"].(string)
	pos++
	// fields (optional)
	fields := []string{}
	if tokens[pos].(map[string]interface{})["Type"] == "(" {
		pos++
		for {
			if tokens[pos].(map[string]interface{})["Type"] == "IDENT" {
				fields = append(fields, tokens[pos].(map[string]interface{})["Value"].(string))
				pos++
				if tokens[pos].(map[string]interface{})["Type"] == "," {
					pos++
					continue
				}
				if tokens[pos].(map[string]interface{})["Type"] == ")" {
					pos++
					break
				}
				return nil, fmt.Errorf("expected ',' or ')' after field name")
			} else {
				return nil, fmt.Errorf("expected field name in field list")
			}
		}
	}
	// VALUES
	if !consume("VALUES") {
		return nil, fmt.Errorf("expected VALUES keyword")
	}
	// values
	if tokens[pos].(map[string]interface{})["Type"] != "(" {
		return nil, fmt.Errorf("expected '(' before values")
	}
	pos++
	values := []interface{}{}
	for {
		tok := tokens[pos].(map[string]interface{})
		if tok["Type"] == "STRING" {
			values = append(values, tok["Value"])
			pos++
		} else if tok["Type"] == "NUMBER" {
			values = append(values, tok["Value"])
			pos++
		} else if tok["Type"] == "IDENT" {
			values = append(values, tok["Value"])
			pos++
		} else if tok["Type"] == "," {
			pos++
			continue
		} else if tok["Type"] == ")" {
			pos++
			break
		} else {
			return nil, fmt.Errorf("unexpected token in values list: %v", tok["Type"])
		}
	}
	return &InsertSQL{Table: table, Fields: fields, Values: values}, nil
}

// ValidateInsertTypes checks if the values match the expected data types in the table schema.
func ValidateInsertTypes(table *model.Table, fields []string, values []interface{}) error {
	if len(fields) != len(values) {
		return fmt.Errorf("field count (%d) does not match value count (%d)", len(fields), len(values))
	}
	for i, fieldName := range fields {
		var fieldSchema *model.Field
		for j := range table.Fields {
			if table.Fields[j].Name == fieldName {
				fieldSchema = &table.Fields[j]
				break
			}
		}
		if fieldSchema == nil {
			return fmt.Errorf("field '%s' not found in table schema", fieldName)
		}
		val := values[i]
		switch fieldSchema.DataType {
		case model.IntType:
			if _, ok := val.(int); !ok {
				if s, ok := val.(string); ok {
					var tmp int
					if _, err := fmt.Sscanf(s, "%d", &tmp); err != nil {
						return fmt.Errorf("field '%s' expects INT, got '%v'", fieldName, val)
					}
				} else {
					return fmt.Errorf("field '%s' expects INT, got '%v'", fieldName, val)
				}
			}
		case model.FloatType:
			if _, ok := val.(float64); !ok {
				if s, ok := val.(string); ok {
					var tmp float64
					if _, err := fmt.Sscanf(s, "%f", &tmp); err != nil {
						return fmt.Errorf("field '%s' expects FLOAT, got '%v'", fieldName, val)
					}
				} else {
					return fmt.Errorf("field '%s' expects FLOAT, got '%v'", fieldName, val)
				}
			}
		case model.BoolType:
			if _, ok := val.(bool); !ok {
				if s, ok := val.(string); ok {
					if s != "true" && s != "false" {
						return fmt.Errorf("field '%s' expects BOOL, got '%v'", fieldName, val)
					}
				} else {
					return fmt.Errorf("field '%s' expects BOOL, got '%v'", fieldName, val)
				}
			}
		case model.VarcharType, model.JSONType:
			if _, ok := val.(string); !ok {
				return fmt.Errorf("field '%s' expects STRING, got '%v'", fieldName, val)
			}
		case model.DateType, model.DateTimeType:
			if _, ok := val.(string); !ok {
				return fmt.Errorf("field '%s' expects DATE/DATETIME string, got '%v'", fieldName, val)
			}
		case model.UUIDType:
			if _, ok := val.(string); !ok {
				return fmt.Errorf("field '%s' expects UUID string, got '%v'", fieldName, val)
			}
		case model.BinaryType:
			if _, ok := val.([]byte); !ok {
				return fmt.Errorf("field '%s' expects binary ([]byte), got '%v'", fieldName, val)
			}
		}
	}
	return nil
}
