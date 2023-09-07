package lexer

import (
	"reflect"
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/stretchr/testify/require"
)

// TestNew verifies that the New function initializes the Lexer correctly.
func TestNew(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	l := New()
	is.Equal(StartPosition(), l.position)
	is.Equal("", l.content)

	nlf1 := reflect.ValueOf(eofLexer)
	nlf2 := reflect.ValueOf(l.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// TestWithPosition validates the position update.
func TestWithPosition(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	newPos := Position{row: 2, col: 4}
	l := New().WithPosition(newPos)
	is.Equal(newPos, l.position)
}

// TestWithContent validates the content update.
func TestWithContent(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	newContent := "\\x.x"
	l := New().WithContent(newContent)
	is.Equal(newContent, l.content)
}

// TestWithNextLexerFunc validates the lexer function update.
func TestWithNextLexerFunc(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	newFunc := func(l Lexer) (monad.Maybe[Token], Lexer) {
		return monad.Some(Token{tokenType: ILLEGAL}), l
	}
	l := New().WithNextLexerFunc(newFunc)
	nlf1 := reflect.ValueOf(newFunc)
	nlf2 := reflect.ValueOf(l.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// TestNext verifies that tokens are generated correctly.
func TestNext(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	// Define a simple lexer function to use in the test
	simpleLexerFunc := func(l Lexer) (monad.Maybe[Token], Lexer) {
		return monad.Some(Token{tokenType: IDENT, literal: "x"}), l
	}

	l := New().WithContent("\\x.x").WithNextLexerFunc(simpleLexerFunc)

	// Call Next() and check the token
	tokenMaybe, _ := l.Next()
	token := tokenMaybe.Value()

	is.Equal(IDENT, token.tokenType)
	is.Equal(Literal("x"), token.Literal())
}
