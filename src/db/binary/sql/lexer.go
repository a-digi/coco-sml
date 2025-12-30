package sql

import "strings"

// LexerToken represents a single token in the SQL input.
type LexerToken struct {
	Type  string
	Value string
}

// Token types for basic SQL SELECT parsing.
const (
	TokenSelect   = "SELECT"
	TokenFrom     = "FROM"
	TokenWhere    = "WHERE"
	TokenAnd      = "AND"
	TokenComma    = ","
	TokenAsterisk = "*"
	TokenIdent    = "IDENT"
	TokenOpEq     = "="
	TokenEOF      = "EOF"
)

// Lexer splits the SQL input into tokens.
// Note: This lexer is basic and only supports simple SELECT queries with = operator.
// It does not handle quoted identifiers, string literals, or other SQL operators.
// For more advanced SQL support, extend this lexer accordingly.
func Lexer(input string) []LexerToken {
	input = strings.TrimSpace(input)
	tokens := []LexerToken{}
	words := strings.Fields(input)
	for _, word := range words {
		upper := strings.ToUpper(word)
		switch upper {
		case "SELECT":
			tokens = append(tokens, LexerToken{Type: TokenSelect, Value: word})
		case "FROM":
			tokens = append(tokens, LexerToken{Type: TokenFrom, Value: word})
		case "WHERE":
			tokens = append(tokens, LexerToken{Type: TokenWhere, Value: word})
		case "AND":
			tokens = append(tokens, LexerToken{Type: TokenAnd, Value: word})
		case "*":
			tokens = append(tokens, LexerToken{Type: TokenAsterisk, Value: word})
		default:
			if strings.Contains(word, ",") {
				for _, part := range strings.Split(word, ",") {
					if part != "" {
						tokens = append(tokens, LexerToken{Type: TokenIdent, Value: part})
					}
					if part != word {
						tokens = append(tokens, LexerToken{Type: TokenComma, Value: ","})
					}
				}
			} else if strings.Contains(word, "=") {
				parts := strings.Split(word, "=")
				if parts[0] != "" {
					tokens = append(tokens, LexerToken{Type: TokenIdent, Value: parts[0]})
				}
				tokens = append(tokens, LexerToken{Type: TokenOpEq, Value: "="})
				if len(parts) > 1 && parts[1] != "" {
					tokens = append(tokens, LexerToken{Type: TokenIdent, Value: parts[1]})
				}
			} else {
				tokens = append(tokens, LexerToken{Type: TokenIdent, Value: word})
			}
		}
	}
	tokens = append(tokens, LexerToken{Type: TokenEOF, Value: ""})

	return tokens
}
