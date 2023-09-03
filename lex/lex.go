// Package lex implements a lexer for lambda c.
// It converts a stream of characters from an io.Reader into a slice of tokens.
//
// The lexer recognizes tokens of several types, such as identifiers,
// assignments, modules, namespace dereferences, lambdas, dots, strings,
// and parentheses.
//
// Usage:
//
// To use the lexer, call the Lex function with an io.Reader as input:
//
//	tokens, err := lex.Lex(reader)
//
// The Lex function returns a slice of tokens or an error.
//
// Token and TokenType:
//
// The lexer creates tokens of type Token. Each token has a type and a value.
// The type is one of the predefined token types (e.g., IDENT, ASSIGN, MODULE, etc.),
// and the value is the string representation of the token.
//
// See the lex_example_test.go file for example usage.
package lex

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// Lex reads an io.Reader input and converts it into a slice of tokens.
// It recognizes several types of tokens like identifiers, assignments,
// modules, namespace dereferences, lambdas, dots, strings, and parentheses.
// It returns an error if it encounters an invalid character or if there is
// an error reading from the input.
func Lex(input io.Reader) ([]Token, error) {
	var tokens []Token
	reader := bufio.NewReader(input)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't read input: %w", err)
		}
		if unicode.IsSpace(r) {
			continue
		}
		switch r {
		case ':':
			exists, err := peekRune(reader, '=')
			if err != nil {
				return nil, fmt.Errorf("can't read input after ':': %w", err)
			}
			if exists {
				tokens = append(tokens, Token{ASSIGN, ":="})
			}
		case '|':
			tokens = append(tokens, Token{MODULE, "|"})
		case '-':
			exists, err := peekRune(reader, '>')
			if err != nil {
				return nil, fmt.Errorf("can't read input after '-': %w", err)
			}
			if exists {
				tokens = append(tokens, Token{NSDEREF, "->"})
			}
		case '\\':
			tokens = append(tokens, Token{LAMBDA, "\\"})
		case '.':
			tokens = append(tokens, Token{DOT, "."})
		case '"':
			if err := reader.UnreadRune(); err != nil {
				return nil, fmt.Errorf("can't unread string rune \": %w", err)
			}
			token, err := parseString(reader)
			if err == io.EOF {
				tokens = append(tokens, token)
				break
			}
			if err != nil {
				return nil, fmt.Errorf("can't parse string: %w", err)
			}
			tokens = append(tokens, token)
		case '(':
			tokens = append(tokens, Token{LPAREN, "("})
		case ')':
			tokens = append(tokens, Token{RPAREN, ")"})
		default:
			if err := reader.UnreadRune(); err != nil {
				return nil, fmt.Errorf("can't unread identifier rune: %w", err)
			}
			token, err := parseIdent(reader)
			if err == io.EOF {
				tokens = append(tokens, token)
				break
			}
			if err != nil {
				return nil, fmt.Errorf("can't parse identifier: %w", err)
			}
			tokens = append(tokens, token)
		}
	}

	tokens = append(tokens, Token{EOF, ""})
	return tokens, nil
}

// isIdentRune checks if a rune is valid for an identifier.
// A valid identifier rune is any graphic character except for
// ':', '|', '-', '\\', '.', '"', '(', and ')'.
func isIdentRune(r rune) bool {
	return unicode.IsGraphic(r) && !unicode.IsSpace(r) &&
		r != ':' && r != '|' && r != '-' && r != '\\' &&
		r != '.' && r != '"' && r != '(' && r != ')'
}

// peekRune reads the next rune from the reader and checks if it is
// the expectedNext rune. It returns a boolean indicating whether the
// next rune is the expectedNext rune and an error if there is one.
// It unreads the rune if it is not the expectedNext rune or if there
// is an error reading it.
func peekRune(reader *bufio.Reader, expectedNext rune) (bool, error) {
	next, _, err := reader.ReadRune()
	if err == io.EOF {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if next == expectedNext {
		return true, nil
	}
	return false, reader.UnreadRune()
}

// parseString reads a string from the reader until it encounters an unescaped
// double quote character ("). It returns the string token and an error if there
// is one. If the end of the input is reached before the closing double quote,
// it returns the string read so far and an io.EOF error.
func parseString(reader *bufio.Reader) (Token, error) {
	// skip first " character
	_, _, err := reader.ReadRune()
	if err != nil {
		return Token{}, err
	}

	var str []rune
	var escaped bool
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			return Token{STRING, string(str)}, err
		}
		if err != nil {
			return Token{}, err
		}
		if escaped {
			str = append(str, r)
			escaped = false
		} else if r == '\\' {
			str = append(str, r)
			escaped = true
		} else if r == '"' {
			break
		} else {
			str = append(str, r)
		}
	}
	return Token{STRING, string(str)}, nil
}

// parseIdent reads an identifier from the reader starting with the given rune.
// It reads runes until it encounters a rune that is not valid for an identifier.
// It returns the identifier token and an error if there is one. If the end of the
// input is reached before completing the identifier, it returns the identifier read
// so far and an io.EOF error.
func parseIdent(reader *bufio.Reader) (Token, error) {
	ident := []rune{}
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return Token{IDENT, string(ident)}, err
		}
		if !isIdentRune(r) {
			err := reader.UnreadRune()
			if err != nil {
				return Token{}, err
			}
			break
		}
		ident = append(ident, r)
	}
	return Token{IDENT, string(ident)}, nil
}
