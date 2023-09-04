package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// dispatcherLexer serves as a pivotal LexerFunc that orchestrates the lexical analysis process.
// It acts as a central dispatcher that delegates specific lexing tasks to other, more specialized, lexers.
// As a result, it encapsulates the mechanism that directs the flow of control based on the type of token
// encountered. It is, in essence, the nucleus of the lexer subsystem, making high-level decisions based
// on a first-pass scan of each rune.
//
// A noteworthy feature of this function is its handling of corner cases:
//  1. If the content to be lexed is empty, it returns an EOF (End-of-File) token, signaling the lexer
//     to cease operation.
//  2. If a newline character is encountered, it returns an EOL (End-of-Line) token and prepares the lexer
//     to commence scanning the new line.
//
// For whitespace characters, the function simply advances the lexer's position without producing a token.
//
// If any of the predefined operator characters are encountered (e.g., '.', '(', ')', '|', ':', '-'),
// the function delegates the task to the operatorLexer for further scrutiny.
//
// If a string delimiter '"' is encountered, the function transitions control to the stringLexer.
//
// For all other characters, it assumes that an identifier or another type of token is being formed
// and delegates this task to the identifierLexer.
//
// The crux of this function is the use of monad.Maybe and setting the NextLexerFunc field of the returned
// Lexer object. By doing so, it elegantly marries functional programming paradigms with the procedural
// nature of lexing, achieving a modular, maintainable, and robust design.
func dispatcherLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	if len(l.content) == 0 {
		return monad.Some(Token{EOF, l.position, ""}), l.WithNextLexerFunc(nil)
	}

	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	if x == '\n' {
		return monad.Some(Token{EOL, l.position, ""}), l.
			WithPosition(l.position.newRow()).
			WithContent(xs).
			WithNextLexerFunc(dispatcherLexer)
	}

	if unicode.IsSpace(x) {
		return monad.None[Token](), l.
			WithPosition(l.position.advanceCol()).
			WithContent(xs).
			WithNextLexerFunc(dispatcherLexer)
	}

	if strings.ContainsAny(string(x), "\\.()|:-") {
		return monad.None[Token](), l.WithNextLexerFunc(operatorLexer)
	}

	if x == '"' {
		return monad.None[Token](), l.WithNextLexerFunc(stringLexer)
	}

	return monad.None[Token](), l.WithNextLexerFunc(identifierLexer)
}
