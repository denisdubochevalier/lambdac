package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLiteralString verifies that the String method returns the correct string representation of a Literal.
func TestLiteralString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    Literal
		expected string
	}{
		{Literal("lambda"), "lambda"},
		{Literal("x"), "x"},
		{Literal(""), ""},
		{Literal("123"), "123"},
		{Literal("!@#$%^&*()"), "!@#$%^&*()"},
		{Literal("whitespace "), "whitespace "},
		{Literal(" "), " "},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.expected, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			is.Equal(testCase.expected, testCase.input.String())
		})
	}
}
