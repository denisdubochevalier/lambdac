package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenTypeString(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	testCases := []struct {
		tokenType TokenType
		expected  string
	}{
		{ILLEGAL, "ILLEGAL"},
		{EOF, "EOF"},
		{EOL, "EOL"},
		{IDENT, "IDENT"},
		{STRING, "STRING"},
		{LAMBDA, "\\"},
		{DOT, "."},
		{LPAREN, "("},
		{RPAREN, ")"},
		{MODULE, "|"},
		{NSDEREF, "->"},
		{ASSIGN, ":="},
		{TokenType(-1), "UNKNOWN"},   // Negative value
		{TokenType(1000), "UNKNOWN"}, // Out-of-bounds value
	}

	for _, testCase := range testCases {
		t.Run(testCase.expected, func(t *testing.T) {
			is.Equal(testCase.expected, testCase.tokenType.String())
		})
	}
}
