package sql

import (
	"errors"
	"fmt"
	"github.com/a-digi/coco-sml/src/db/binary/model"
)

// SQLQuery represents a parsed SQL query (SELECT, INSERT, UPDATE, DELETE).
type SQLQuery struct {
	Table      string
	Fields     []string
	Conditions map[string]string // field -> value (simple equality only)
	Limit      int               // Optional LIMIT clause
}

var ErrInvalidSQL = errors.New("invalid SQL query")

// ParseSQL parses a basic SELECT SQL query using the lexer and builds a SQLQuery struct.
func ParseSQL(query string) (*SQLQuery, error) {
	tokens := Lexer(query)
	pos := 0
	consume := func(expected string) bool {
		if pos < len(tokens) && tokens[pos].Type == expected {
			pos++
			return true
		}
		return false
	}
	// SELECT
	if !consume(TokenSelect) {
		return nil, ErrInvalidSQL
	}

	// fields
	fields := []string{}
	for {
		if tokens[pos].Type == TokenAsterisk {
			fields = append(fields, "*")
			pos++
			break
		}
		if tokens[pos].Type == TokenIdent {
			fields = append(fields, tokens[pos].Value)
			pos++
			if tokens[pos].Type == TokenComma {
				pos++
				continue
			}
			break
		} else {
			break
		}
	}
	// FROM
	if !consume(TokenFrom) {
		return nil, ErrInvalidSQL
	}
	// table name
	if tokens[pos].Type != TokenIdent {
		return nil, ErrInvalidSQL
	}
	table := tokens[pos].Value
	pos++
	// WHERE (optional)
	conditions := map[string]string{}
	if consume(TokenWhere) {
		for {
			if tokens[pos].Type != TokenIdent {
				break
			}
			field := tokens[pos].Value
			pos++
			if !consume(TokenOpEq) {
				return nil, ErrInvalidSQL
			}
			if tokens[pos].Type != TokenIdent {
				return nil, ErrInvalidSQL
			}
			value := tokens[pos].Value
			pos++
			conditions[field] = value
			if consume(TokenAnd) {
				continue
			}
			break
		}
	}
	// LIMIT (optional)
	limit := 0
	if consume(TokenLimit) {
		if tokens[pos].Type == TokenIdent {
			parsed := 0
			fmt.Sscanf(tokens[pos].Value, "%d", &parsed)
			limit = parsed
			pos++
		}
	}

	return &SQLQuery{Table: table, Fields: fields, Conditions: conditions, Limit: limit}, nil
}

// TableSQL represents a parsed CREATE TABLE SQL statement.
type TableSQL struct {
	TableName string
	Fields    []TableField
}

type TableField struct {
	Name string
	Type string
}

// ParseTableSQL parses a basic CREATE TABLE SQL statement using the lexer and builds a model.Table struct.
// Example: CREATE TABLE users (id INT, name STRING, email STRING)
func ParseTableSQL(query string) (*model.Table, error) {
	tokens := Lexer(query)
	pos := 0
	consume := func(expected string) bool {
		if pos < len(tokens) && tokens[pos].Type == expected {
			pos++
			return true
		}
		return false
	}
	// CREATE
	if !consume(TokenCreate) {
		return nil, errors.New("expected CREATE keyword")
	}
	// TABLE
	if !consume(TokenTable) {
		return nil, errors.New("expected TABLE keyword")
	}
	// table name
	if tokens[pos].Type != TokenIdent {
		return nil, errors.New("expected table name")
	}
	tableName := tokens[pos].Value
	pos++
	// (
	if !consume(TokenLParen) {
		return nil, errors.New("expected '('")
	}
	fields := []model.Field{}
	for {
		if tokens[pos].Type == TokenRParen {
			pos++
			break
		}
		if tokens[pos].Type != TokenIdent {
			return nil, errors.New("expected field name")
		}
		fieldName := tokens[pos].Value
		pos++
		if tokens[pos].Type != TokenIdent {
			return nil, errors.New("expected field type")
		}
		fieldType := tokens[pos].Value
		pos++
		var dt model.DataType
		switch fieldType {
		case "INT", "int":
			dt = model.IntType
		case "FLOAT", "float":
			dt = model.FloatType
		case "BOOL", "bool":
			dt = model.BoolType
		case "DATE", "date":
			dt = model.DateType
		case "DATETIME", "datetime":
			dt = model.DateTimeType
		case "BINARY", "binary":
			dt = model.BinaryType
		case "UUID", "uuid":
			dt = model.UUIDType
		case "JSON", "json":
			dt = model.JSONType
		case "VARCHAR", "varchar":
			dt = model.VarcharType
		default:
			return nil, fmt.Errorf("unsupported field type: %s", fieldType)
		}
		// NOT NULL Unterstützung
		nullable := true
		if tokens[pos].Type == TokenIdent && tokens[pos].Value == "NOT" {
			pos++
			if tokens[pos].Type == TokenIdent && tokens[pos].Value == "NULL" {
				nullable = false
				pos++
			}
		}
		fields = append(fields, model.Field{
			Name:     fieldName,
			DataType: dt,
			Nullable: nullable,
		})
		if tokens[pos].Type == TokenComma {
			pos++
			continue
		}
	}
	// optional semicolon
	if tokens[pos].Type == TokenSemicolon {
		pos++
	}
	return &model.Table{
		Name:        tableName,
		Description: "Created via SQL",
		Fields:      fields,
		EntryCount:  0,
	}, nil
}
