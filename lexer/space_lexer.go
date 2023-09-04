package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// spaceLexer is responsible for handling whitespace characters other than EOL.
// This function is invoked if neither EOF nor EOL conditions are met.
//
// Parameters:
//   - l: The current Lexer object.
//
// Returns:
//   - monad.Maybe[Token]: A Maybe monad encapsulating a token, if any.
//     None is returned if the character is a space.
//   - Lexer: A new Lexer object with a potentially modified state.
//
// Upon encountering a space, it advances the column position and
// continues the lexer process by setting the next lexer function to eofLexer.
//
// If a space is not encountered, it inspects the character to determine
// which lexer function should be called next, for example, operatorLexer or stringLexer.
func spaceLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	if unicode.IsSpace(x) {
		return monad.None[Token](), l.
			WithPosition(l.position.advanceCol()).
			WithContent(xs).
			WithNextLexerFunc(eofLexer)
	}

	if strings.ContainsAny(string(x), "\\.()|:-") {
		return operatorLexer(l)
	}

	if x == '"' {
		return stringLexer(l)
	}

	return identifierLexer(l)
}
