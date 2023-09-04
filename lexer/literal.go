package lexer

// Literal is a type used to store the value of a token.
// It is a simple wrapper around the string type.
type Literal string

// String returns the string representation of the Literal.
// This method is used to implement the fmt.Stringer interface.
func (l Literal) String() string {
	return string(l)
}
