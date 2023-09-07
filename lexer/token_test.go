package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Example-based tests for Token methods
func TestTokenMethods(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	// Initialize Token
	token := Token{
		tokenType: IDENT,
		position:  Position{row: 1, col: 5},
		literal:   "foobar",
	}

	// Test Type method
	is.Equal(IDENT, token.Type())

	// Test Position method
	is.Equal(Position{row: 1, col: 5}, token.Position())

	// Test Literal method
	is.Equal("foobar", string(token.Literal()))
}
