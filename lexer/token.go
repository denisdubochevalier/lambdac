package lexer

// Token represents a lexeme or a sequence of characters that have a collective meaning.
// It contains the type of the token (e.g. IDENT, ASSIGN, etc.), the position
// in the input where the token starts, and the literal value of the token.
type Token struct {
	tokenType TokenType
	position  Position
	literal   Literal
}
