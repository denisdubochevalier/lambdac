package lexer

import (
	"strings"
	"testing"
	"unicode"

	"github.com/denisdubochevalier/monad"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/require"
)

func TestSpaceLexerExampleBased(t *testing.T) {
	t.Parallel()

	// Define example-based test cases
	testCases := []struct {
		name         string
		content      string
		expectedType TokenType
	}{
		{"SpaceChar", " ", ILLEGAL},
		{"TabChar", "\t", ILLEGAL},
		{"OperatorChar", "\\", LAMBDA},
		{"StringChar", "\"\"", STRING},
	}

	// Execute example-based tests
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			lexer := New().WithContent(testCase.content)
			maybeToken, _ := spaceLexer(lexer)

			if testCase.expectedType == ILLEGAL {
				_, ok := maybeToken.(monad.Nothing[Token])
				is.True(ok)
			} else {
				tokenJust, ok := maybeToken.(monad.Just[Token])
				is.True(ok)
				is.Equal(testCase.expectedType, tokenJust.Value().tokenType)
			}
		})
	}
}

func TestSpaceLexerPropertyBased(t *testing.T) {
	t.Parallel()

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 1000
	properties := gopter.NewProperties(parameters)

	// Define a property for the spaceLexer function
	properties.Property("spaceLexer handles runes correctly", prop.ForAll(
		func(r rune) bool {
			content := string(r)
			lexer := New().WithContent(content)
			maybeToken, _ := spaceLexer(lexer)

			if unicode.IsSpace(r) {
				_, ok := maybeToken.(monad.Nothing[Token])
				return ok
			} else if strings.ContainsAny(content, "\\.()|:-") {
				tokenJust, ok := maybeToken.(monad.Just[Token])
				return ok && tokenJust.Value().tokenType != ILLEGAL
			}
			return true
		},
		gen.Rune(),
	))

	properties.TestingRun(t)
}
