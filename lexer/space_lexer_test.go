package lexer

import (
	"strings"
	"testing"
	"unicode"

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
				is.True(maybeToken.Nothing())
			} else {
				is.True(maybeToken.Just())
				is.Equal(testCase.expectedType, maybeToken.Value().tokenType)
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
				return maybeToken.Nothing()
			} else if strings.ContainsAny(content, "\\.()|:-") {
				return maybeToken.Just() && maybeToken.Value().tokenType != ILLEGAL
			}
			return true
		},
		gen.Rune(),
	))

	properties.TestingRun(t)
}
