package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// stringLexer is a specialized LexerFunc designed for tokenizing strings in a lambda calculus language.
// It utilizes a form of recursive descent to build up the string token incrementally. Each invocation
// of the function peels off one character from the front of the remaining input and decides on the next
// course of action based on that character and the state accumulated so far.
//
// This function faces unique challenges, such as handling escaped quotes within strings, and capturing
// illegal states like unclosed strings or strings containing newlines.
//
// Recursive descent is typically used in parsing rather than lexing, but in this unique context, it serves
// to modularize and isolate the logic required for string tokenization.
//
// The function uses a local higher-order function `strLex` that encapsulates the core recursion logic. The
// `strLex` function is tail-recursive and operates on an accumulating monad.Either[Token] to build up the token's
// Literal value. It offers early termination through monad.NewRVal in case of illegal states, thereby folding
// both the success and failure states into a unified approach.
//
// Finally, the function returns a monad.Maybe[Token] encapsulating the resulting token if a legal string is
// found or the illegal state otherwise. It also returns a new Lexer with an updated state to be used in
// subsequent lexing operations.
func stringLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	// Invalid string with EOF after a single "
	if len(l.content) < 1 {
		return monad.Some(Token{ILLEGAL, l.position, Literal(l.content)}), l.WithNextLexerFunc(nil)
	}

	// Skip the first '"' character and start the recursion
	_, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	val := recursiveStringLexer(Token{STRING, l.position, ""}, xs)
	content, _ := strings.CutPrefix(l.content, string(val.Value().literal))

	return monad.Some(val.Value()), l.
		WithPosition(l.position.advanceColBy(utf8.RuneCountInString(val.Value().literal.String()))).
		WithContent(content).
		WithNextLexerFunc(eofLexer)
}

// handleEscapeCharacter handles the escape character '\' within a string literal.
func handleEscapeCharacter(xs string) (string, rune, bool) {
	x2, _ := utf8.DecodeRuneInString(xs)
	switch x2 {
	case '"':
		return xs[1:], x2, false
	default:
		return xs[1:], x2, true
	}
}

// appendStringToken appends the current character to the existing string token.
func appendStringToken(t Token, x rune) Token {
	t.literal = t.literal + Literal(string(x))
	return t
}

// recursiveStringLexer is the recursive function to handle string lexing.
func recursiveStringLexer(t Token, xs string) monad.Either[Token] {
	if len(xs) == 0 {
		return monad.NewRVal(Token{ILLEGAL, t.Position(), ""})
	}

	x, size := utf8.DecodeRuneInString(xs)
	xs = xs[size:]

	if x == '\n' {
		return monad.NewRVal(Token{ILLEGAL, t.Position(), ""})
	}

	if x == '"' {
		return monad.NewRVal(Token{STRING, t.Position(), t.Literal()})
	}

	if x == '\\' {
		keepBackslash := false
		xs, x, keepBackslash = handleEscapeCharacter(xs)
		if keepBackslash {
			t = appendStringToken(t, '\\')
		}
	}

	return recursiveStringLexer(appendStringToken(t, x), xs)
}
