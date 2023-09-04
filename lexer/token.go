package lexer

// Token represents a lexeme or a sequence of characters that have a collective meaning.
// It contains the type of the token (e.g. IDENT, ASSIGN, etc.), the position
// in the input where the token starts, and the literal value of the token.
type Token struct {
	tokenType TokenType
	position  Position
	literal   Literal
}

// Type returns the TokenType of a token instance.
// This provides an immutable way to access the type, upholding
// the principle of encapsulation.
func (t Token) Type() TokenType {
	return t.tokenType
}

// Position returns the position of a token instance, similarly
// to TokenType, protecting it from mutations and ensuring
// data integrity.
func (t Token) Position() Position {
	return t.position
}

// Literal gets the token literal, ie. the actual string in the
// source text. It protects the underlying value from mutations.
func (t Token) Literal() Literal {
	return t.literal
}
