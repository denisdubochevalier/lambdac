package lexer

import (
	"reflect"
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/require"
)

// Property-Based Testing
func TestCompositeLexerProperties(t *testing.T) {
	t.Parallel()

	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property(
		"compositeLexer returns a monad.Maybe[Token] and updated Lexer",
		prop.ForAll(
			func(a string) bool {
				lexer := New().WithContent(a)
				result, _ := lexer.Next()
				return result != nil
			},
			gen.AlphaString(),
		),
	)

	properties.TestingRun(t)
}

// Referential Transparency Check
func TestCompositeLexerReferentialTransparency(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	lexer1 := New().WithContent(":").WithNextLexerFunc(compositeLexer)
	lexer2 := New().WithContent(":").WithNextLexerFunc(compositeLexer)

	result1, _ := lexer1.Next()
	result2, _ := lexer2.Next()

	is.Equal(result1, result2, "compositeLexer fails referential transparency check")
}

// More specialized tests
func TestCompositeLexerSpecialCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input         string
		expected      monad.Maybe[Token]
		nextLexerFunc lexerFunc
	}{
		{":=", monad.Some(Token{ASSIGN, StartPosition(), ":="}), eofLexer},
		{"->", monad.Some(Token{NSDEREF, StartPosition(), "->"}), eofLexer},
		{":", monad.None[Token](), identifierLexer},
		{":>", monad.None[Token](), identifierLexer},
		{"-=", monad.None[Token](), identifierLexer},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.input, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			lexer := New().WithContent(testCase.input).WithNextLexerFunc(compositeLexer)
			result, lexer := lexer.Next()

			is.Equal(testCase.expected, result, "Unexpected result for input: "+testCase.input)
			// Test next lexer func returned
			nlf1 := reflect.ValueOf(testCase.nextLexerFunc)
			nlf2 := reflect.ValueOf(lexer.nextLexerFunc)
			is.Equal(
				nlf1.Pointer(),
				nlf2.Pointer(),
				"Unexpected next lexer func for input: "+testCase.input,
			)
		})
	}
}
