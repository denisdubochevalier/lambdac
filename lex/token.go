package lex

// TokenType is the type of a parsed element
type TokenType int

// Token represent a parsed element
type Token struct {
	Type  TokenType
	Value string
}

// Define the token types
const (
	ILLEGAL TokenType = iota
	EOF
	WS

	IDENT   // m, recursive_fact, fact, ...
	ASSIGN  // :=
	MODULE  // |
	NSDEREF // ->
	LAMBDA  // \
	DOT     // .
	STRING  // "github.com/foo/bar", "text/lexer", ...
	LPAREN  // (
	RPAREN  // )
)

// String implements fmt.Stringer
func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "illegal"
	case EOF:
		return "EOF"
	case WS:
		return "WS"
	case IDENT:
		return "IDENT"
	case ASSIGN:
		return "ASSIGN"
	case MODULE:
		return "MODULE"
	case NSDEREF:
		return "NSDEREF"
	case LAMBDA:
		return "LAMBDA"
	case DOT:
		return "DOT"
	case STRING:
		return "STRING"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	}
	return ""
}
