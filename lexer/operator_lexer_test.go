package lexer

import (
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/stretchr/testify/require"
)

func TestOperatorLexer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		content      string
		expectedType TokenType
		expectedLit  Literal
	}{
		{"LAMBDA", "\\", LAMBDA, "\\"},
		{"DOT", ".", DOT, "."},
		{"LPAREN", "(", LPAREN, "("},
		{"RPAREN", ")", RPAREN, ")"},
		{"MODULE", "|", MODULE, "|"},
		{"Handling to compositLexer: ASSIGN", ":=", ASSIGN, ":="},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			// Initialize lexer and position
			lexer := New().WithContent(testCase.content)

			// Invoke the operatorLexer function
			maybeToken, newLexer := operatorLexer(lexer)

			if testCase.expectedType == ILLEGAL {
				_, ok := maybeToken.(monad.Nothing[Token])
				is.True(ok)
			} else {
				_, ok := maybeToken.(monad.Just[Token])
				is.True(ok)
				token := maybeToken.Value()
				is.Equal(testCase.expectedType, token.tokenType)
				is.Equal(testCase.expectedLit, token.literal)
			}

			// Ensure that the new lexer's position and content have been updated correctly
			is.Equal(StartPosition().advanceCol(), newLexer.position)
		})
	}
}
