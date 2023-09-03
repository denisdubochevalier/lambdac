package lex

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIdentRune(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title  string
		input  rune
		output bool
	}{
		{
			title:  "latin letestCaseer",
			input:  'a',
			output: true,
		},
		{
			title:  "greek letestCaseer",
			input:  'λ',
			output: true,
		},
		{
			title:  "accent",
			input:  '`',
			output: true,
		},
		{
			title:  "emoji",
			input:  '👍',
			output: true,
		},
		{
			title:  "space",
			input:  ' ',
			output: false,
		},
		{
			title:  "tabulation",
			input:  '\t',
			output: false,
		},
		{
			title:  "carriot return",
			input:  '\r',
			output: false,
		},
		{
			title:  "line feed",
			input:  '\n',
			output: false,
		},
		{
			title:  "colon",
			input:  ':',
			output: false,
		},
		{
			title:  "pipe",
			input:  '|',
			output: false,
		},
		{
			title:  "dash",
			input:  '-',
			output: false,
		},
		{
			title:  "antislash",
			input:  '\\',
			output: false,
		},
		{
			title:  "dot",
			input:  '.',
			output: false,
		},
		{
			title:  "quote",
			input:  '"',
			output: false,
		},
		{
			title:  "lparen",
			input:  '(',
			output: false,
		},
		{
			title:  "rparen",
			input:  ')',
			output: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			is.Equal(testCase.output, isIdentRune(testCase.input))
		})
	}
}

func TestPeekRune(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		title          string
		input          string
		expectedNext   rune
		expectedExists bool
	}{
		{
			title:          "exists",
			input:          ":=",
			expectedNext:   '=',
			expectedExists: true,
		},
		{
			title:          "does not exist",
			input:          ":",
			expectedNext:   '=',
			expectedExists: false,
		},
		{
			title:          "empty input",
			input:          "",
			expectedNext:   '=',
			expectedExists: false,
		},
		{
			title:          "other character",
			input:          ":a",
			expectedNext:   '=',
			expectedExists: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			reader := bufio.NewReader(strings.NewReader(testCase.input))
			reader.ReadRune() // consume first rune

			exists, err := peekRune(reader, testCase.expectedNext)
			is.NoError(err)
			is.Equal(testCase.expectedExists, exists)
		})
	}
}

func TestParseString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title  string
		input  string
		output string
		err    error
	}{
		{
			title:  "simple string",
			input:  "\"hello\"",
			output: "hello",
			err:    nil,
		},
		{
			title:  "escaped quote",
			input:  "\"hello\\\"\"",
			output: "hello\\\"",
			err:    nil,
		},
		{
			title:  "unicode characters",
			input:  "\"hello 世界\"",
			output: "hello 世界",
			err:    nil,
		},
		{
			title:  "new line",
			input:  "\"hello\\nworld\"",
			output: "hello\\nworld",
			err:    nil,
		},
		{
			title:  "empty string",
			input:  "\"\"",
			output: "",
			err:    nil,
		},
		{
			title:  "EOF before closing quote",
			input:  "\"hello",
			output: "hello",
			err:    io.EOF,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			reader := bufio.NewReader(strings.NewReader(testCase.input))

			result, err := parseString(reader)
			if err != io.EOF {
				is.NoError(err)
			}
			is.Equal(STRING, result.Type)
			is.Equal(testCase.output, result.Value)
		})
	}
}

func TestParseIdent(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title  string
		input  string
		output string
		err    error
	}{
		{
			title:  "simple identifier",
			input:  "hello",
			output: "hello",
			err:    io.EOF,
		},
		{
			title:  "identifier with numbers",
			input:  "var123",
			output: "var123",
			err:    io.EOF,
		},
		{
			title:  "identifier followed by space",
			input:  "foo ",
			output: "foo",
			err:    nil,
		},
		{
			title:  "identifier followed by special character",
			input:  "bar(",
			output: "bar",
			err:    nil,
		},
		{
			title:  "empty input",
			input:  "",
			output: "",
			err:    io.EOF,
		},
		{
			title:  "input with special character only",
			input:  "(",
			output: "",
			err:    nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			reader := bufio.NewReader(strings.NewReader(testCase.input))

			result, err := parseIdent(reader)
			if err != io.EOF {
				is.NoError(err)
			}
			is.Equal(IDENT, result.Type)
			is.Equal(testCase.output, result.Value)
		})
	}
}

func TestLex(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		title  string
		input  string
		output []Token
	}{
		{
			title: "multiple tokens",
			input: "hello := | -> \\ . ( )",
			output: []Token{
				{IDENT, "hello"},
				{ASSIGN, ":="},
				{MODULE, "|"},
				{NSDEREF, "->"},
				{LAMBDA, "\\"},
				{DOT, "."},
				{LPAREN, "("},
				{RPAREN, ")"},
				{EOF, ""},
			},
		},
		{
			title: "assignment without lexpr",
			input: " := | -> \\ . ( ",
			output: []Token{
				{ASSIGN, ":="},
				{MODULE, "|"},
				{NSDEREF, "->"},
				{LAMBDA, "\\"},
				{DOT, "."},
				{LPAREN, "("},
				{EOF, ""},
			},
		},
		{
			title:  "single identifier",
			input:  "foo",
			output: []Token{{IDENT, "foo"}, {EOF, ""}},
		},
		{
			title:  "EOF before closing parenthesis",
			input:  "(foo",
			output: []Token{{LPAREN, "("}, {IDENT, "foo"}, {EOF, ""}},
		},
		{
			title:  "single identifier followed by space",
			input:  "hello ",
			output: []Token{{IDENT, "hello"}, {EOF, ""}},
		},
		{
			title:  "unterminated string",
			input:  "\"foo",
			output: []Token{{STRING, "foo"}, {EOF, ""}},
		},
		{
			title:  "lambda",
			input:  "\\",
			output: []Token{{LAMBDA, "\\"}, {EOF, ""}},
		},
		{
			title:  "string",
			input:  "\"hello\"",
			output: []Token{{STRING, "hello"}, {EOF, ""}},
		},
		{
			title:  "empty code",
			input:  "",
			output: []Token{{EOF, ""}},
		},
		{
			title:  "assignment without rexpr",
			input:  " hello:=",
			output: []Token{{IDENT, "hello"}, {ASSIGN, ":="}, {EOF, ""}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			reader := strings.NewReader(testCase.input)

			tokens, err := Lex(reader)
			is.NoError(err)

			is.Equal(testCase.output, tokens)
		})
	}
}
