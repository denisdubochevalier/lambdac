package lexer

import (
	"reflect"
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/stretchr/testify/require"
)

// Testing eolLexer when the content starts with an EOL character
func TestEolLexerWithEolCharacter(t *testing.T) {
	// t.Parallel()
	is := require.New(t)

	lexer := New().WithContent("\n").WithNextLexerFunc(eolLexer)
	result, updatedLexer := lexer.Next()

	// Asserting that an EOL Token is returned
	is.Equal(monad.Some(Token{EOL, StartPosition(), ""}), result)

	// Asserting that the lexer position has advanced to a new row
	is.Equal(StartPosition().newRow().row, updatedLexer.position.row)

	// Asserting the nextLexerFunc is eofLexer
	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(updatedLexer.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// Testing eolLexer when the content starts with multiple EOL character
func TestEolLexerWithMultipleEolCharacter(t *testing.T) {
	// t.Parallel()
	is := require.New(t)

	lexer := New().WithContent("\n\n\n").WithNextLexerFunc(eolLexer)
	result, updatedLexer := lexer.Next()

	// Asserting that an EOL Token is returned
	is.Equal(monad.Some(Token{EOL, StartPosition(), ""}), result)

	// Asserting that the lexer position has advanced to a new row
	is.Equal(StartPosition().newRow().newRow().newRow().row, updatedLexer.position.row)

	// Asserting the nextLexerFunc is eofLexer
	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(updatedLexer.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// Testing eolLexer when the content starts with EOL character and contains further EOL character later
func TestEolLexerWithEolCharacterAfterNonEolCharacter(t *testing.T) {
	// t.Parallel()
	is := require.New(t)

	lexer := New().WithContent("\n \n").WithNextLexerFunc(eolLexer)
	result, updatedLexer := lexer.Next()

	// Asserting that an EOL Token is returned
	is.Equal(monad.Some(Token{EOL, StartPosition(), ""}), result)

	// Asserting that the lexer position has advanced to a new row
	is.Equal(StartPosition().newRow().row, updatedLexer.position.row)

	// Asserting the nextLexerFunc is eofLexer
	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(updatedLexer.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// Testing eolLexer when the content does not start with an EOL character
func TestEolLexerWithNonEolCharacter(t *testing.T) {
	// t.Parallel()
	is := require.New(t)

	lexer := New().WithContent(" ").WithNextLexerFunc(eolLexer)
	result, updatedLexer := lexer.Next()

	// No EOL Token should be returned; this should delegate to spaceLexer
	is.Equal(monad.None[Token](), result)

	// Asserting that the lexer position has advanced to a new row
	is.Equal(StartPosition().row, updatedLexer.position.row)

	// Asserting the nextLexerFunc has switched to spaceLexer
	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(updatedLexer.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}
