package lexer

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/require"
)

func TestStringLexer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input  string
		output Token
	}{
		{"", Token{ILLEGAL, StartPosition(), ""}},
		{"\"Hello World\"", Token{STRING, StartPosition(), "Hello World"}},
		{"\"Hello\\\"World\"", Token{STRING, StartPosition(), `Hello"World`}},
		{"\"Hello\\World\"", Token{STRING, StartPosition(), "Hello\\World"}},

		{"\"Unclosed String", Token{ILLEGAL, StartPosition(), ""}},

		{"\"StringWithNewLine\n\"", Token{ILLEGAL, StartPosition(), ""}},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.input, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			lexer := New().WithContent(testCase.input)
			tokenMaybe, _ := stringLexer(lexer)
			token := tokenMaybe.Value()
			is.Equal(testCase.output, token)
		})
	}
}

func TestStringLexerProperty(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("stringLexer should return STRING tokens for valid inputs", prop.ForAll(
		func(str string) bool {
			// Wrap in quotes to simulate a string literal
			input := "\"" + str + "\""
			lexer := New().WithContent(input)
			tokenMaybe, _ := stringLexer(lexer)
			token := tokenMaybe.Value()

			return token.Type() == STRING && token.Literal() == Literal(str)
		},
		gen.AlphaString(),
	))

	properties.TestingRun(t)
}
