package sql

import (
	"errors"
	"fmt"
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
