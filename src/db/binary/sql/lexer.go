package sql

import (
	"strings"
	"unicode"
)

// LexerToken represents a single token in the SQL input.
type LexerToken struct {
	Type  string
	Value string
}

// Token types for SQL parsing.
const (
	TokenCreate    = "CREATE"
	TokenTable     = "TABLE"
	TokenSelect    = "SELECT"
	TokenInsert    = "INSERT"
	TokenInto      = "INTO"
	TokenValues    = "VALUES"
	TokenFrom      = "FROM"
	TokenWhere     = "WHERE"
	TokenAnd       = "AND"
	TokenComma     = ","
	TokenAsterisk  = "*"
	TokenIdent     = "IDENT"
	TokenOpEq      = "="
	TokenOpNeq     = "!="
	TokenOpLt      = "<"
	TokenOpGt      = ">"
	TokenOpLe      = "<="
	TokenOpGe      = ">="
	TokenLParen    = "("
	TokenRParen    = ")"
	TokenSemicolon = ";"
	TokenString    = "STRING"
	TokenNumber    = "NUMBER"
	TokenEOF       = "EOF"
	TokenUnknown   = "UNKNOWN"
	TokenLimit     = "LIMIT"
)

// Lexer splits the SQL input into tokens.
// Now supports: =, !=, <, >, <=, >=, (, ), ;, string literals, and robust comma handling.
func Lexer(input string) []LexerToken {
	input = strings.TrimSpace(input)
	tokens := []LexerToken{}
	runes := []rune(input)
	length := len(runes)
	pos := 0

	next := func() rune {
		if pos < length {
			r := runes[pos]
			pos++
			return r
		}
		return 0
	}

	peek := func() rune {
		if pos < length {
			return runes[pos]
		}
		return 0
	}

	skipWhitespace := func() {
		for unicode.IsSpace(peek()) {
			next()
		}
	}

	for pos < length {
		skipWhitespace()
		r := peek()
		if r == 0 {
			break
		}
		// Handle single-char tokens
		switch r {
		case ',':
			next()
			tokens = append(tokens, LexerToken{Type: TokenComma, Value: ","})
			continue
		case '*':
			next()
			tokens = append(tokens, LexerToken{Type: TokenAsterisk, Value: "*"})
			continue
		case '(':
			next()
			tokens = append(tokens, LexerToken{Type: TokenLParen, Value: "("})
			continue
		case ')':
			next()
			tokens = append(tokens, LexerToken{Type: TokenRParen, Value: ")"})
			continue
		case ';':
			next()
			tokens = append(tokens, LexerToken{Type: TokenSemicolon, Value: ";"})
			continue
		case '\'', '"': // String literal
			quote := next()
			start := pos
			for peek() != 0 && peek() != quote {
				next()
			}
			val := string(runes[start:pos])
			if peek() == quote {
				next()
			}
			tokens = append(tokens, LexerToken{Type: TokenString, Value: val})
			continue
		}
		// Handle numbers (integer/float)
		if unicode.IsDigit(r) {
			start := pos
			next()
			for unicode.IsDigit(peek()) || peek() == '.' {
				next()
			}
			val := string(runes[start:pos])
			tokens = append(tokens, LexerToken{Type: TokenNumber, Value: val})
			continue
		}
		// Handle operators
		if r == '=' {
			next()
			tokens = append(tokens, LexerToken{Type: TokenOpEq, Value: "="})
			continue
		}
		if r == '!' && peekN(runes, pos, length, 1) == '=' {
			next(); next()
			tokens = append(tokens, LexerToken{Type: TokenOpNeq, Value: "!="})
			continue
		}
		if r == '<' {
			next()
			if peek() == '=' {
				next()
				tokens = append(tokens, LexerToken{Type: TokenOpLe, Value: "<="})
			} else {
				tokens = append(tokens, LexerToken{Type: TokenOpLt, Value: "<"})
			}
			continue
		}
		if r == '>' {
			next()
			if peek() == '=' {
				next()
				tokens = append(tokens, LexerToken{Type: TokenOpGe, Value: ">="})
			} else {
				tokens = append(tokens, LexerToken{Type: TokenOpGt, Value: ">"})
			}
			continue
		}
		// Identifiers and keywords
		if unicode.IsLetter(r) || r == '_' {
			start := pos
			next()
			for unicode.IsLetter(peek()) || unicode.IsDigit(peek()) || peek() == '_' {
				next()
			}
			val := string(runes[start:pos])
			upper := strings.ToUpper(val)
			switch upper {
			case "CREATE":
				tokens = append(tokens, LexerToken{Type: TokenCreate, Value: val})
			case "TABLE":
				tokens = append(tokens, LexerToken{Type: TokenTable, Value: val})
			case "SELECT":
				tokens = append(tokens, LexerToken{Type: TokenSelect, Value: val})
			case "INSERT":
				tokens = append(tokens, LexerToken{Type: TokenInsert, Value: val})
			case "INTO":
				tokens = append(tokens, LexerToken{Type: TokenInto, Value: val})
			case "VALUES":
				tokens = append(tokens, LexerToken{Type: TokenValues, Value: val})
			case "FROM":
				tokens = append(tokens, LexerToken{Type: TokenFrom, Value: val})
			case "WHERE":
				tokens = append(tokens, LexerToken{Type: TokenWhere, Value: val})
			case "AND":
				tokens = append(tokens, LexerToken{Type: TokenAnd, Value: val})
			case "LIMIT":
				tokens = append(tokens, LexerToken{Type: TokenLimit, Value: val})
			default:
				tokens = append(tokens, LexerToken{Type: TokenIdent, Value: val})
			}
			continue
		}
		// Unknown token
		tokens = append(tokens, LexerToken{Type: TokenUnknown, Value: string(r)})
		next()
	}
	tokens = append(tokens, LexerToken{Type: TokenEOF, Value: ""})
	return tokens
}

// peekN peeks ahead n runes in the input.
func peekN(runes []rune, pos, length, n int) rune {
	if pos+n < length {
		return runes[pos+n]
	}
	return 0
}
