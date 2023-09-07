package lexer

import (
	"reflect"
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/stretchr/testify/require"
)

// Testing eofLexer when the content is empty
func TestEofLexerWithEmptyContent(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	lexer := New().WithContent("").WithNextLexerFunc(eofLexer)
	result, updatedLexer := lexer.Next()

	// Asserting the token is an EOF
	is.Equal(monad.Some(Token{EOF, StartPosition(), ""}), result)

	// Asserting the nextLexerFunc is nil, indicating termination
	is.Nil(updatedLexer.nextLexerFunc)
}

// Testing eofLexer when the content is not empty
func TestEofLexerWithNonEmptyContent(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	lexer := New().WithContent("\n").WithNextLexerFunc(eofLexer)
	result, updatedLexer := lexer.Next()

	// No EOF should be returned; this should delegate to eolLexer
	is.Equal(monad.Some(Token{EOL, StartPosition(), ""}), result)

	// Asserting the nextLexerFunc has switched to eolLexer
	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(updatedLexer.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}
